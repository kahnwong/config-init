/*
Copyright © 2025 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var composeOptions = []string{
	"app",
	"caddy",
	"clickhouse",
	"elasticsearch",
	"mongo",
	"postgres",
}

var composeDestMapping = map[string]string{
	"app":           "compose.yaml",
	"caddy":         "compose-caddy.yaml",
	"clickhouse":    "compose-clickhouse.yaml",
	"elasticsearch": "compose-elasticsearch.yaml",
	"mongo":         "compose-mongo.yaml",
	"postgres":      "compose-db.yaml",
}

var composeCmd = &cobra.Command{
	Use:       "compose",
	Short:     "Init compose.yaml",
	ValidArgs: composeOptions,
	Run: func(cmd *cobra.Command, args []string) {
		requireTemplateOption(args)

		option := args[0]
		filename := fmt.Sprintf("compose-%s.yaml", option)
		destFile := composeDestMapping[option]

		template.WriteConfig("compose", filename, destFile)
	},
}

func init() {
	rootCmd.AddCommand(composeCmd)
}
