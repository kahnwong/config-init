/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"path/filepath"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var githubActionsOptions = []string{
	"cloudflare-pages",
	"docker-build",
	"github-pages",
	"go-test",
	"goreleaser",
	"nix",
	"node-test",
	"pre-commit",
	"rust-test",
}

var githubActionsMapping = map[string][2]string{
	"cloudflare-pages": {"cloudflare-pages.yaml", "deploy.yaml"},
	"docker-build":     {"docker-build.yaml", "build.yaml"},
	"github-pages":     {"github-pages.yaml", "deploy.yaml"},
	"go-test":          {"go-test.yaml", "go-test.yaml"},
	"goreleaser":       {"goreleaser.yaml", "release.yaml"},
	"nix":              {"nix.yaml", "deploy.yaml"},
	"node-test":        {"node-test.yaml", "node-test.yaml"},
	"pre-commit":       {"pre-commit.yaml", "pre-commit.yaml"},
	"python-test":      {"python-test.yaml", "python-test.yaml"},
	"rust-test":        {"rust-test.yaml", "rust-test.yaml"},
}

var githubActionsCmd = &cobra.Command{
	Use:       "github-actions",
	Short:     "Init Github Actions configs",
	ValidArgs: githubActionsOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		option := args[0]
		filename, destFile := mapOptionSeparate(option, githubActionsMapping)

		template.CreateDir(filepath.Join(".github", "workflows"))
		template.WriteConfig("github-actions", filename, filepath.Join(".github", "workflows", destFile))
	},
}

func init() {
	rootCmd.AddCommand(githubActionsCmd)
}
