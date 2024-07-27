/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
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
}

var githubActionsCmd = &cobra.Command{
	Use:       "github-actions",
	Short:     "Init Github Actions config",
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
		}

		_ = template.CreateDir(filepath.Join(".github", "workflows"))
		template.WriteConfig("github-actions", filename, filepath.Join(".github", "workflows", destFile))
	},
}

func init() {
	rootCmd.AddCommand(githubActionsCmd)
}
