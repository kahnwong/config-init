/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Init Makefile",
	Run: func(cmd *cobra.Command, args []string) {
		template.WriteConfig("make", "Makefile", "Makefile")
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
