package cmd

import (
	"os"

	cliBase "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/rs/zerolog/log"
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

		keyPath, err := cliBase.ExpandHome(sourceFile)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to expand home path for %s", sourceFile)
		}
		contentBytes, err := os.ReadFile(keyPath)
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
