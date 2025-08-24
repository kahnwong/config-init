/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var renovateCmd = &cobra.Command{
	Use:   "renovate",
	Short: "Init renovate ci config",
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("renovate", "renovate.json", "renovate.json")
	},
}

func init() {
	rootCmd.AddCommand(renovateCmd)
}
