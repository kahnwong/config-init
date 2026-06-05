/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"path/filepath"

	"github.com/kahnwong/config-init/template"

	"github.com/spf13/cobra"
)

var dependabotMapping = [][2]string{
	{"dependabot.yaml", "dependabot.yaml"},
}

var dependabotCmd = &cobra.Command{
	Use:   "dependabot",
	Short: "Init dependabot ci config",
	Run: func(cmd *cobra.Command, args []string) {
		template.CreateDir(".github")
		for _, mapping := range dependabotMapping {
			template.WriteConfig("dependabot", mapping[0], filepath.Join(".github", mapping[1]))
		}
	},
}

func init() {
	rootCmd.AddCommand(dependabotCmd)
}
