/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var preCommitCmd = &cobra.Command{
	Use:   "pre-commit",
	Short: "Init pre-commit",
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("pre-commit", "pre-commit-config.yaml", ".pre-commit-config.yaml")
	},
}

func init() {
	rootCmd.AddCommand(preCommitCmd)
}
