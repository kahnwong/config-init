package cmd

import (
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var direnvCmd = &cobra.Command{
	Use:   "direnv",
	Short: "Init direnv",
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("direnv", "envrc", ".envrc")
		template.ExecCommand("direnv", "allow")
	},
}

func init() {
	rootCmd.AddCommand(direnvCmd)
}
