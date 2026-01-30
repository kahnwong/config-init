/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var dockerfileOptions = []string{
	"binary",
	"go",
	"nix",
	"node",
	"python",
	"rust",
	"spa",
	"static-site",
}

var dockerfileCmd = &cobra.Command{
	Use:       "dockerfile",
	Short:     "Init Dockerfile",
	ValidArgs: dockerfileOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		filename := fmt.Sprintf("%s.Dockerfile", args[0])
		template.WriteConfig("dockerfile", filename, "Dockerfile")
	},
}

func init() {
	rootCmd.AddCommand(dockerfileCmd)
}
