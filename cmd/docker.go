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

var dockerOptions = []string{
	"binary",
	"fastapi",
	"go",
	"nix",
}

var dockerCmd = &cobra.Command{
	Use:       "docker",
	Short:     "Init Dockerfile",
	ValidArgs: dockerOptions,
	Run: func(cmd *cobra.Command, args []string) {
		// validate args
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		template.WriteConfig("docker", args[0], "Dockerfile")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
