/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"log/slog"
	"os"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "Init Rust",
}

var rustInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Rust package",
	Run: func(cmd *cobra.Command, args []string) {
		stdout, err := cli_base.ExecCommand("cargo", "init")
		if err != nil {
			slog.Error("failed to run cargo init", "err", err)
			os.Exit(1)
		}
		if stdout != "" {
			slog.Info("cargo init", "stdout", stdout)
		}

		template.WriteConfig("rust", "rust_toolchain.toml", "rust_toolchain.toml")
	},
}

func init() {
	rustCmd.AddCommand(rustInitCmd)
	rootCmd.AddCommand(rustCmd)
}
