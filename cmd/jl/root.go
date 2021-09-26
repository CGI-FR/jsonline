package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	over "github.com/Trendyol/overlog"

	"github.com/adrienaury/go-template/pkg/jsonline"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type globalFlags struct {
	verbosity string
	debug     bool
	jsonlog   bool
	colormode string
}

type templateFlags struct {
	columns  []string
	template string
	filename string
}

type RootCommand struct {
	cobra.Command
}

func NewRootCommand() (*RootCommand, error) {
	// nolint: exhaustivestruct
	rootCmd := cobra.Command{
		Use:     fmt.Sprintf("%v", name),
		Short:   "JSONLine templating",
		Long:    `Order keys and format output of JSON lines.`,
		Args:    cobra.NoArgs,
		Run:     run,
		Version: fmt.Sprintf("%v (commit=%v date=%v by=%v)", version, commit, buildDate, builtBy),
		Example: "" +
			fmt.Sprintf(`  %s -c '{name: first, type: string}' -c '{name: second, type: string}' <dirty.jsonl`, name) + "\n" +
			fmt.Sprintf(`  %s -t '{"first":"string","second":"string"}' <dirty.jsonl`, name),
	}

	cobra.OnInitialize(initConfig)

	gf := globalFlags{
		verbosity: "error",
		debug:     false,
		jsonlog:   false,
		colormode: "auto",
	}

	tf := templateFlags{
		columns:  nil,
		template: "",
		filename: "./row.yml",
	}

	rootCmd.PersistentFlags().StringVarP(&gf.verbosity, "verbosity", "v", gf.verbosity,
		"set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&gf.debug, "debug", gf.debug, "add debug information to logs (very slow)")
	rootCmd.PersistentFlags().BoolVar(&gf.jsonlog, "log-json", gf.jsonlog, "output logs in JSON format")
	rootCmd.PersistentFlags().StringVar(&gf.colormode, "color", gf.colormode,
		"use colors in log outputs : yes, no or auto")

	rootCmd.PersistentFlags().SortFlags = false

	rootCmd.Flags().StringArrayVarP(&tf.columns, "column", "c", tf.columns,
		`inline column definition in minified YAML (-c {name: title, type: string})`+"\n"+
			`use this flag multiple times, one for each column`+"\n"+
			`possible types : string, numeric, boolean, binary, datetime, time, timestamp, row, auto, hidden`)
	rootCmd.Flags().StringVarP(&tf.template, "template", "t", tf.template,
		`row template definition in JSON (-t {"title":"string"})`+"\n"+
			`possible types : string, numeric, boolean, binary, datetime, time, timestamp, auto, hidden`)
	rootCmd.Flags().StringVarP(&tf.filename, "filename", "f", tf.filename, "name of row template filename")

	rootCmd.Flags().SortFlags = false

	if err := bindViper(rootCmd); err != nil {
		return nil, err
	}

	// rootCmd.AddCommand(<package>.NewCommand(rootCmd.CommandPath()))

	return &RootCommand{rootCmd}, nil
}

func bindViper(rootCmd cobra.Command) error {
	err := viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = viper.BindPFlag("log_json", rootCmd.PersistentFlags().Lookup("log-json"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = viper.BindPFlag("color", rootCmd.PersistentFlags().Lookup("color"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func initConfig() {
	initViper()

	color := computeColor()

	jsonlog := viper.GetBool("log_json")
	if jsonlog {
		log.Logger = zerolog.New(os.Stderr)
	} else {
		// nolint: exhaustivestruct
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: !color})
	}

	debug := viper.GetBool("debug")
	if debug {
		log.Logger = log.Logger.With().Caller().Logger()
	}

	over.New(log.Logger)

	verbosity := viper.GetString("verbosity")
	switch verbosity {
	case "trace", "5":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Debug().Msg("logger level set to trace")
	case "debug", "4":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("logger level set to debug")
	case "info", "3":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn", "2":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error", "1":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}

func computeColor() bool {
	color := false

	switch strings.ToLower(viper.GetString("color")) {
	case "auto":
		if isatty.IsTerminal(os.Stdout.Fd()) && runtime.GOOS != "windows" {
			color = true
		}
	case "yes", "true", "1", "on", "enable":
		color = true
	}

	return color
}

func initViper() {
	viper.SetEnvPrefix("jl") // will be uppercased automatically
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.jl")

	if err := viper.ReadInConfig(); err != nil {
		// nolint: errorlint
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

func run(cmd *cobra.Command, args []string) {
	t, err := createTemplate(cmd)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse row template")
		os.Exit(1)
	}

	ri := NewJSONRowIterator(os.Stdin)

	defer ri.Close()

	over.AddGlobalFields("line-number")

	startTime := time.Now()

	for i := 1; ri.Next(); i++ {
		over.MDC().Set("line-number", i)

		row := ri.Value()
		fmt.Println(t.Create(row))
	}

	duration := time.Since(startTime)

	log.Info().Int("return", 0).Stringer("duration", duration).Msg("end of process")
}

func getTemplateFlags(cmd *cobra.Command) (*templateFlags, error) {
	tf := &templateFlags{
		columns:  []string{},
		template: "",
		filename: "",
	}

	var err error

	tf.filename, err = cmd.Flags().GetString("filename")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	tf.columns, err = cmd.Flags().GetStringArray("column")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	tf.template, err = cmd.Flags().GetString("template")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if len(tf.columns) > 0 && len(tf.template) > 0 {
		return nil, ErrForbiddenTemplateAndColumnFlags
	}

	log.Debug().
		Str("filename", tf.filename).
		Str("template", tf.template).
		Str("columns", fmt.Sprintf("%v", tf.columns)).
		Msg("template flags")

	return tf, nil
}

func createTemplate(cmd *cobra.Command) (jsonline.Template, error) {
	tf, err := getTemplateFlags(cmd)
	if err != nil {
		return nil, err
	}

	columns := []ColumnDefinition{}

	for _, columnDef := range tf.columns {
		colDef := ColumnDefinition{Name: "", Type: "", Columns: nil}

		err = yaml.Unmarshal([]byte(columnDef), &colDef)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		columns = append(columns, colDef)
	}

	t, err := ParseRowDefinition(tf.filename)
	if err != nil {
		return nil, err
	}

	if len(columns) > 0 {
		t, err = parse(jsonline.NewTemplate(), columns)
		if err != nil {
			return nil, err
		}
	} else if len(tf.template) > 0 {
		t, err = createTemplateFromString(tf.template)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
