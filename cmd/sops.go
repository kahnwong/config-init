package cmd

import (
	"os"

	cliBase "github.com/kahnwong/cli-base"
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

		contentBytes, err := os.ReadFile(cliBase.ExpandHome(sourceFile))
		if err != nil {
			panic(err)
		}

		template.WriteFile(".sops.yaml", contentBytes, 0664)
		template.ExecCommand("git", "add", ".sops.yaml")
	},
}

func init() {
	rootCmd.AddCommand(sopsCmd)
}
