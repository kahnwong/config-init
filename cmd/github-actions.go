/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"os"
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

var githubActionsCmd = &cobra.Command{
	Use:       "github-actions",
	Short:     "Init Github Actions configs",
	ValidArgs: githubActionsOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		filename := args[0]
		var destFile string
		switch filename {
		case "cloudflare-pages":
			filename = "cloudflare-pages.yaml"
			destFile = "deploy.yaml"
		case "docker-build":
			filename = "docker-build.yaml"
			destFile = "build.yaml"
		case "github-pages":
			filename = "github-pages.yaml"
			destFile = "deploy.yaml"
		case "go-test":
			filename = "go-test.yaml"
			destFile = "go-test.yaml"
		case "goreleaser":
			filename = "goreleaser.yaml"
			destFile = "release.yaml"
		case "nix":
			filename = "nix.yaml"
			destFile = "deploy.yaml"
		case "node-test":
			filename = "node-test.yaml"
			destFile = "node-test.yaml"
		case "pre-commit":
			filename = "pre-commit.yaml"
			destFile = "pre-commit.yaml"
		case "python-test":
			filename = "python-test.yaml"
			destFile = "python-test.yaml"
		case "rust-test":
			filename = "rust-test.yaml"
			destFile = "rust-test.yaml"
		}

		template.CreateDir(filepath.Join(".github", "workflows"))
		template.WriteConfig("github-actions", filename, filepath.Join(".github", "workflows", destFile))
	},
}

func init() {
	rootCmd.AddCommand(githubActionsCmd)
}
