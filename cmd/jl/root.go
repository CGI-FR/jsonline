// Copyright (C) 2021 CGI France
//
// This file is part of JL.
//
// JL is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// JL is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with JL.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	over "github.com/Trendyol/overlog"

	"github.com/cgi-fr/jsonline/pkg/jsonline"
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
}

type templateFlags struct {
	template string
	filename string
}

type RootCommand struct {
	cobra.Command
}

func NewRootCommand() (*RootCommand, error) {
	rootCmd := cobra.Command{ //nolint:exhaustivestruct
		Use:     fmt.Sprintf("%v", name),
		Short:   "JSONLine templating",
		Long:    `Order keys and enforce format of JSON lines.`,
		Args:    cobra.NoArgs,
		Run:     run,
		Version: fmt.Sprintf("%v (commit=%v date=%v by=%v)", version, commit, buildDate, builtBy),
		Example: "" +
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
		template: "{}",
		filename: "./row.yml",
	}

	rootCmd.PersistentFlags().StringVarP(&gf.verbosity, "verbosity", "v", gf.verbosity,
		"set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5)")
	rootCmd.PersistentFlags().BoolVar(&gf.debug, "debug", gf.debug, "add debug information to logs (very slow)")
	rootCmd.PersistentFlags().BoolVar(&gf.jsonlog, "log-json", gf.jsonlog, "output logs in JSON format")
	rootCmd.PersistentFlags().StringVar(&gf.colormode, "color", gf.colormode,
		"use colors in log outputs : yes, no or auto")

	rootCmd.PersistentFlags().SortFlags = false

	//nolint:lll
	rootCmd.Flags().StringVarP(&tf.template, "template", "t", tf.template,
		`row template definition (-t {"name":"format"} or -t {"name":"format(type)"}) or -t {"name":"format(type):format"})`+"\n"+
			`possible formats : string, numeric, boolean, binary, datetime, time, timestamp, auto, hidden`+"\n"+
			`possible types : int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, byte, rune, string, []byte, time.Time, json.Number`)
	rootCmd.Flags().StringVarP(&tf.filename, "filename", "f", tf.filename, "name of row template filename")

	rootCmd.Flags().SortFlags = false

	if err := bindViper(rootCmd); err != nil {
		return nil, err
	}

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
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: !color}) //nolint:exhaustivestruct
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
	ti, to, err := createTemplate(cmd)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse row template")
		os.Exit(1)
	}

	importer := ti.GetImporter(os.Stdin)
	exporter := to.GetExporter(os.Stdout)
	streamer := jsonline.NewStreamer(importer, exporter)

	over.AddGlobalFields("line-number")

	startTime := time.Now()
	i := 0
	p := func(row jsonline.Row, err error) error {
		over.MDC().Set("line-number", i)
		i++

		if err != nil {
			log.Error().Err(err).Msg("failed to process JSON line")
		}

		if row != nil {
			log.Trace().Str("raw", row.DebugString()).RawJSON("export", []byte(row.String())).Msg("processed JSON line")
		} else {
			log.Trace().Str("raw", "<nil>").Str("export", "<nil>").Msg("processed JSON line")
		}

		return nil
	}

	if err := streamer.WithProcessor(p).Stream(); err != nil {
		log.Error().Err(err).Msg("streamer failed")
	}

	duration := time.Since(startTime)

	log.Info().Int("return", 0).Stringer("duration", duration).Msg("end of process")
}

func getTemplateFlags(cmd *cobra.Command) (*templateFlags, error) {
	tf := &templateFlags{
		template: "",
		filename: "",
	}

	var err error

	tf.filename, err = cmd.Flags().GetString("filename")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	tf.template, err = cmd.Flags().GetString("template")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	logFlags(tf)

	return tf, nil
}

func logFlags(tf *templateFlags) {
	var js json.RawMessage

	tmp := log.Debug().
		Str("filename", tf.filename)

	if json.Unmarshal([]byte(tf.template), &js) != nil {
		tmp = tmp.Str("in", tf.template)
	} else {
		tmp = tmp.RawJSON("in", []byte(tf.template))
	}

	tmp.Msg("template flags")
}

func createTemplate(cmd *cobra.Command) (jsonline.Template, jsonline.Template, error) {
	tf, err := getTemplateFlags(cmd)
	if err != nil {
		return nil, nil, err
	}

	ti, to, err := ParseRowDefinition(tf.filename)
	if err != nil {
		return nil, nil, err
	}

	if len(tf.template) > 0 && tf.template != "{}" {
		ti, to, err = createTemplateFromString(tf.template)
		if err != nil {
			return nil, nil, err
		}
	}

	return ti, to, nil
}
