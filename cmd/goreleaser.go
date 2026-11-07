/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var goreleaserOptions = []string{
	"go",
	"rust",
}

var goreleaserMapping = map[string][2]string{
	"go":   {"goreleaser.yaml", ".goreleaser.yaml"},
	"rust": {"goreleaser-rust.yaml", ".goreleaser.yaml"},
}

var goreleaserCmd = &cobra.Command{
	Use:       "goreleaser",
	Short:     "Init goreleaser config",
	ValidArgs: goreleaserOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		option := args[0]
		filename, destFile := mapOptionSeparate(option, goreleaserMapping)

		template.WriteConfig("goreleaser", filename, destFile)
	},
}

func init() {
	rootCmd.AddCommand(goreleaserCmd)
}
