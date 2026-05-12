package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var sopsAccounts = []string{
	"personal",
	"work",
}
var sopsFileMapping = map[string]string{
	"personal": "~/.sops.yaml",
	"work":     "~/.sops-work.yaml",
}

var sopsCmd = &cobra.Command{
	Use:       "sops",
	Short:     "Init sops",
	ValidArgs: sopsAccounts,
	Run: func(cmd *cobra.Command, args []string) {
		requireArgs(args, errMsgSopsAccount)

		account := args[0]
		sourceFile := sopsFileMapping[account]

		keyPath, err := cli_base.ExpandHome(sourceFile)
		if err != nil {
			slog.Error("failed to expand home path", "source_file", sourceFile, "err", err)
			os.Exit(1)
		}
		contentBytes, err := os.ReadFile(keyPath)
		if err != nil {
			panic(err)
		}

		template.WriteFile(".sops.yaml", contentBytes, 0664)
		stdout, err := cli_base.ExecCommand("git", "add", ".sops.yaml")
		if err != nil {
			slog.Error("failed to add .sops.yaml to git", "stdout", stdout, "err", err)
			os.Exit(1)
		}
		fmt.Println(stdout)
	},
}

func init() {
	rootCmd.AddCommand(sopsCmd)
}
