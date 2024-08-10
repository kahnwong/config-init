/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var preCommitOptions = []string{
	"markdownlint",
	"tfdocs",
}

var preCommitCmd = &cobra.Command{
	Use:       "pre-commit",
	Short:     "Init pre-commit",
	ValidArgs: preCommitOptions,
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("pre-commit", "pre-commit-config.yaml", ".pre-commit-config.yaml")

		// init pre-commit
		template.ExecCommand("git", "init")
		template.ExecCommand("git", "add", ".pre-commit-config.yaml")
		template.ExecCommand("pre-commit", "install")

		// hooks configurations
		if len(args) > 0 {
			hookConfigOption := args[0]
			var filename string
			var destFile string

			switch hookConfigOption {
			case "markdownlint":
				filename = "markdownlint.yaml"
				destFile = ".markdownlint.yaml"
			}

			template.WriteConfig("pre-commit", filename, destFile)
			template.ExecCommand("git", "add", destFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(preCommitCmd)
}
