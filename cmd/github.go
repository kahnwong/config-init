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

var githubOptions = []string{
	"dependabot",
}

var githubCmd = &cobra.Command{
	Use:       "github",
	Short:     "Init GitHub-related configs",
	ValidArgs: githubOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		filename := args[0]
		var destFile string
		switch filename {
		case "dependabot":
			filename = "dependabot.yaml"
			destFile = "dependabot.yaml"
		}

		template.CreateDir(".github")
		template.WriteConfig("github", filename, filepath.Join(".github", destFile))
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
}
