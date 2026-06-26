/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/kahnwong/config-init/template"
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

// mapOptionSeparate returns separate filename and destFile from a map
func mapOptionSeparate(option string, mapping map[string][2]string) (filename string, destFile string) {
	if val, ok := mapping[option]; ok {
		return val[0], val[1]
	}
	return option, option
}

// writeConfigAndGitAdd writes a config file and adds it to git
func writeConfigAndGitAdd(templateDir, filename, destFile string) {
	template.WriteConfig(templateDir, filename, destFile)
	stdout, err := cli_base.ExecCommand("git", "add", destFile)
	if err != nil {
		slog.Error("failed to add file to git", "file", destFile, "err", err)
		os.Exit(1)
	}
	fmt.Println(stdout)
}
