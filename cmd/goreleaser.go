/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var goreleaserOptions = []string{
	"linux",
	"darwin",
}

var goreleaserCmd = &cobra.Command{
	Use:       "goreleaser",
	Short:     "Init goreleaser config",
	ValidArgs: goreleaserOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		filename := args[0]
		var destFile string
		switch filename {
		case "linux":
			filename = "goreleaser-linux.yaml"
			destFile = ".goreleaser-linux.yaml"
		case "darwin":
			filename = "goreleaser-darwin.yaml"
			destFile = ".goreleaser-darwin.yaml"
		}

		template.WriteConfig("goreleaser", filename, destFile)
	},
}

func init() {
	rootCmd.AddCommand(goreleaserCmd)
}
