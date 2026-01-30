/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var preCommitOptions = []string{
	"markdownlint",
	"rumdl",
	"tfdocs",
}

var preCommitHookMapping = map[string][2]string{
	"markdownlint": {"markdownlint.yaml", ".markdownlint.yaml"},
	"rumdl":        {"rumdl.toml", ".rumdl.toml"},
	"tfdocs":       {"terraform-docs.yaml", ".terraform-docs.yaml"},
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
			filename, destFile := mapOptionSeparate(hookConfigOption, preCommitHookMapping)

			writeConfigAndGitAdd("pre-commit", filename, destFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(preCommitCmd)
}
