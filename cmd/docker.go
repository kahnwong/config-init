/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kahnwong/config-init/utils"
	"github.com/spf13/cobra"
)

var dockerOptions = []string{
	"nix",
}

var dockerCmd = &cobra.Command{
	Use:       "docker",
	Short:     "Init Dockerfile",
	ValidArgs: dockerOptions,
	Run: func(cmd *cobra.Command, args []string) {
		// validate args
		if len(args) == 0 {
			fmt.Println("Please specify a template option")
			os.Exit(1)
		}

		// main
		option := args[0]
		filename := fmt.Sprintf("%s.Dockerfile", option)
		destFile := "Dockerfile"

		// -- set dest path --
		wd, _ := os.Getwd()
		destPath := filepath.Join(wd, destFile)

		// -- write template --
		template, _ := templatesFS.ReadFile(fmt.Sprintf("templates/docker/%s", filename))
		utils.WriteFile(destPath, string(template), 0664)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
