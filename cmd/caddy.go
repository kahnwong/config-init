/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"

	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var caddyOptions = []string{
	"spa",
}

var caddyCmd = &cobra.Command{
	Use:       "caddy",
	Short:     "Init Caddyfile",
	ValidArgs: caddyOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		filename := fmt.Sprintf("%s.%s", args[0], "Caddyfile")
		template.WriteConfig("caddy", filename, "Caddyfile")
	},
}

func init() {
	rootCmd.AddCommand(caddyCmd)
}
