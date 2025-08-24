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

var caddyOptions = []string{
	"spa",
}

var caddyCmd = &cobra.Command{
	Use:       "caddy",
	Short:     "Init Caddyfile",
	ValidArgs: caddyOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		filename := fmt.Sprintf("%s.%s", args[0], "Caddyfile")
		template.WriteConfig("caddy", filename, "Caddyfile")
	},
}

func init() {
	rootCmd.AddCommand(caddyCmd)
}
