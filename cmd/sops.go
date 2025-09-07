package cmd

import (
	"fmt"
	"os"

	cliBase "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var sopsAccounts = []string{
	"personal",
	"work",
}
var sopsCmd = &cobra.Command{
	Use:       "sops",
	Short:     "Init sops",
	ValidArgs: sopsAccounts,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a sops account")
			os.Exit(1)
		}

		var contentBytes []byte
		var err error

		account := args[0]
		switch account {
		case "personal":
			contentBytes, err = os.ReadFile(cliBase.ExpandHome("~/.sops.yaml"))
			if err != nil {
				panic(err)
			}

		case "work":
			contentBytes, err = os.ReadFile(cliBase.ExpandHome("~/.sops-work.yaml"))
			if err != nil {
				panic(err)
			}
		}

		template.WriteFile(".sops.yaml", contentBytes, 0664)
		template.ExecCommand("git", "add", ".sops.yaml")
	},
}

func init() {
	rootCmd.AddCommand(sopsCmd)
}
