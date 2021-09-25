package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type globalFlags struct {
	verbosity string
	debug     bool
	jsonlog   bool
	colormode string
	template  map[string]string
	filename  string
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
		Example: fmt.Sprintf(`  %s name:string surname:string age:numeric <dirty.jsonl`, name),
	}

	cobra.OnInitialize(initConfig)

	gf := globalFlags{
		verbosity: "error",
		debug:     false,
		jsonlog:   false,
		colormode: "auto",
		template:  map[string]string{},
		filename:  "./row.yml",
	}

	rootCmd.PersistentFlags().StringVarP(&gf.verbosity, "verbosity", "v", gf.verbosity,
		"set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&gf.debug, "debug", gf.debug, "add debug information to logs (very slow)")
	rootCmd.PersistentFlags().BoolVar(&gf.jsonlog, "log-json", gf.jsonlog, "output logs in JSON format")
	rootCmd.PersistentFlags().StringVar(&gf.colormode, "color", gf.colormode,
		"use colors in log outputs : yes, no or auto")
	rootCmd.PersistentFlags().StringToStringVarP(&gf.template, "template", "t", nil,
		`inline template definition (<name>=<type>,<name>=<type>,...)`+"\n"+
			`possible types : string, numeric, boolean, binary, datetime, time, timestamp, row, auto, hidden`)
	rootCmd.PersistentFlags().StringVarP(&gf.filename, "filename", "f", gf.filename, "name of row template filename")

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

	verbosity := viper.GetString("verbosity")
	switch verbosity {
	case "trace", "5":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Debug().Msg("Logger level set to trace")
	case "debug", "4":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Logger level set to debug")
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
	filename, err := cmd.Flags().GetString("filename")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read flag filename")
		os.Exit(1)
	}

	t, err := ParseRowDefinition(filename)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse template file")
		os.Exit(1)
	}

	def, err := cmd.Flags().GetStringToString("template")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read flag template")
		os.Exit(1)
	}

	rowdef := &RowDefinition{Columns: []ColumnDefinition{}}
	for colname, coltype := range def {
		rowdef.Columns = append(rowdef.Columns, ColumnDefinition{Name: colname, Type: coltype, Columns: nil})
	}

	t, err = parse(t, rowdef.Columns)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse flag template")
		os.Exit(1)
	}

	ri := NewJSONRowIterator(os.Stdin)

	defer ri.Close()

	for ri.Next() {
		row := ri.Value()
		fmt.Println(t.Create(row))
	}
}
