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

var composeOptions = []string{
	"app",
	"clickhouse",
	"postgres",
}

var composeCmd = &cobra.Command{
	Use:       "compose",
	Short:     "Init compose.yaml",
	ValidArgs: composeOptions,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		option := args[0]
		filename := fmt.Sprintf("compose-%s.yaml", option)

		var destFile string
		switch option {
		case "app":
			destFile = "compose.yaml"
		case "clickhouse":
			destFile = "compose-clickhouse.yaml"
		case "postgres":
			destFile = "compose-db.yaml"
		}

		template.WriteConfig("compose", filename, destFile)
	},
}

func init() {
	rootCmd.AddCommand(composeCmd)
}
