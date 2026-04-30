/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var dockerOptions = []string{
	"bake",
}

var dockerCmd = &cobra.Command{
	Use:       "docker",
	Short:     "Init Docker",
	ValidArgs: dockerOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		template.WriteConfig("docker", "docker-bake.hcl", "docker-bake.hcl")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
