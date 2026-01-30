/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kahnwong/config-init/template"
	"github.com/spf13/cobra"
)

const (
	errMsgTemplateOption = "Please specify a template option"
	errMsgSopsAccount    = "Please specify a sops account"
)

// requireArgs validates that at least one argument is provided
func requireArgs(args []string, message string) {
	if len(args) == 0 {
		fmt.Println(message)
		os.Exit(1)
	}
}

// requireTemplateOption validates that a template option argument is provided
func requireTemplateOption(args []string) {
	requireArgs(args, errMsgTemplateOption)
}

// mapOption returns the mapped value for a given option from a map
func mapOption(option string, mapping map[string]string) (filename string, destFile string) {
	if val, ok := mapping[option]; ok {
		return val, val
	}
	return option, option
}

// mapOptionSeparate returns separate filename and destFile from a map
func mapOptionSeparate(option string, mapping map[string][2]string) (filename string, destFile string) {
	if val, ok := mapping[option]; ok {
		return val[0], val[1]
	}
	return option, option
}

// createSimpleConfigCommand creates a cobra command that writes a single config file
func createSimpleConfigCommand(use, short, templateDir, filename, destFile string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			template.WriteConfig(templateDir, filename, destFile)
		},
	}
}

// writeConfigAndGitAdd writes a config file and adds it to git
func writeConfigAndGitAdd(templateDir, filename, destFile string) {
	template.WriteConfig(templateDir, filename, destFile)
	template.ExecCommand("git", "add", destFile)
}
