package cmd

import (
	"log/slog"
	"os"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var direnvCmd = &cobra.Command{
	Use:   "direnv",
	Short: "Init direnv",
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("direnv", "envrc", ".envrc")
		stdout, err := cli_base.ExecCommand("direnv", "allow")
		if err != nil {
			slog.Error("failed to allow direnv", "err", err)
			os.Exit(1)
		}
		slog.Info("direnv allow", "stdout", stdout)
	},
}

func init() {
	rootCmd.AddCommand(direnvCmd)
}
