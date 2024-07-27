/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"path/filepath"

	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var dependabotCmd = &cobra.Command{
	Use:   "dependabot",
	Short: "Init dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		_ = template.CreateDir(".github")
		template.WriteConfig("dependabot", "dependabot.yaml", filepath.Join(".github", "dependabot.yaml"))
	},
}

func init() {
	rootCmd.AddCommand(dependabotCmd)
}
