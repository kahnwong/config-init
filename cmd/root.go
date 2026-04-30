/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "config-init",
	Version: version,
	Short:   "Utils to init various configs",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	slog.SetDefault(slog.New(slogzerolog.Option{Logger: &logger}.NewZerologHandler()))

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
