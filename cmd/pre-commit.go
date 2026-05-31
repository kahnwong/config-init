/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var preCommitLatestRevision = "e4f4ade"

var preCommitOptions = []string{
	"markdownlint",
	"rumdl",
	"tfdocs",
	"typos",

	// utils
	"bump", // for bumping pre-commit version specified in existing config
}

var preCommitHookMapping = map[string][2]string{
	"markdownlint": {"markdownlint.yaml", ".markdownlint.yaml"},
	"rumdl":        {"rumdl.toml", ".rumdl.toml"},
	"tfdocs":       {"terraform-docs.yaml", ".terraform-docs.yaml"},
	"typos":        {"typos.toml", "_typos.toml"},
}

var preCommitCmd = &cobra.Command{
	Use:       "pre-commit",
	Short:     "Init pre-commit",
	ValidArgs: preCommitOptions,
	Run: func(cmd *cobra.Command, args []string) {
		// set mode
		mode := "setup"
		if len(args) > 0 {
			if args[0] == "bump" {
				mode = "bump"
			}
		}

		// scenarios
		switch mode {
		case "setup":
			template.WriteConfig("pre-commit", "pre-commit-config.yaml", ".pre-commit-config.yaml")

			// init pre-commit
			stdout, err := cli_base.ExecCommand("bash", "-c", "git init && git add .pre-commit-config.yaml && pre-commit install")
			if err != nil {
				slog.Error("failed to init pre-commit", "err", err)
				os.Exit(1)
			}
			fmt.Println(stdout)

			// hooks configurations
			if len(args) > 0 {
				hookConfigOption := args[0]
				filename, destFile := mapOptionSeparate(hookConfigOption, preCommitHookMapping)

				writeConfigAndGitAdd("pre-commit", filename, destFile)
			}
		case "bump":
			stdout, err := cli_base.ExecCommand("bash", "-c", fmt.Sprintf("yq e -i '.repos[1].rev = \"%s\"' .pre-commit-config.yaml", preCommitLatestRevision))
			if err != nil {
				slog.Error("failed to bump pre-commit", "err", err)
				os.Exit(1)
			}
			fmt.Println(stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(preCommitCmd)
}
