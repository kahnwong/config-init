/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var goreleaserOptions = []string{
	"linux",
	"darwin",
}

var goreleaserMapping = map[string][2]string{
	"linux":  {"goreleaser-linux.yaml", ".goreleaser-linux.yaml"},
	"darwin": {"goreleaser-darwin.yaml", ".goreleaser-darwin.yaml"},
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
