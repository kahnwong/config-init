package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/spf13/cobra"
)

var dvcCmd = &cobra.Command{
	Use:   "dvc",
	Short: "Init dvc",
	Run: func(cmd *cobra.Command, args []string) {
		stdout, err := cli_base.ExecCommand("dvc", "init")
		if err != nil {
			slog.Error("failed to init dvc", "err", err)
			os.Exit(1)
		}
		fmt.Println(stdout)

		writeConfigAndGitAdd("dvc", "config", filepath.Join(".dvc", "config"))
	},
}

func init() {
	rootCmd.AddCommand(dvcCmd)
}
