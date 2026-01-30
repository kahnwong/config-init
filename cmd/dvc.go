package cmd

import (
	"path/filepath"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

var dvcCmd = &cobra.Command{
	Use:   "dvc",
	Short: "Init dvc",
	Run: func(cmd *cobra.Command, args []string) {
		template.ExecCommand("dvc", "init")
		writeConfigAndGitAdd("dvc", "config", filepath.Join(".dvc", "config"))
	},
}

func init() {
	rootCmd.AddCommand(dvcCmd)
}
