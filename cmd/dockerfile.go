/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var dockerfileOptions = []string{
	"binary",
	"fastapi",
	"go",
	"nix",
	"node",
	"spa",
	"static-site",
}

var dockerfileCmd = &cobra.Command{
	Use:       "dockerfile",
	Short:     "Init Dockerfile",
	ValidArgs: dockerfileOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		filename := fmt.Sprintf("%s.%s", args[0], "Dockerfile")
		template.WriteConfig("dockerfile", filename, "Dockerfile")
	},
}

func init() {
	rootCmd.AddCommand(dockerfileCmd)
}
