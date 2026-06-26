/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"log/slog"
	"os"

	cli_base "github.com/kahnwong/cli-base"
	"github.com/spf13/cobra"
)

const pyprojectUvConfig = `
# [tool.uv.build-backend]
# module-root = "api"

# [tool.uv]
# package = false
`

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Init Python",
}

var pythonInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Python package with uv",
	Run: func(cmd *cobra.Command, args []string) {
		runPythonInitCommand("uv", "python", "pin", "3.14")
		runPythonInitCommand("uv", "init", "--package", "--author-from", "git", "--no-description")
		runPythonInitCommand("uv", "venv")

		file, err := os.OpenFile("pyproject.toml", os.O_APPEND|os.O_WRONLY, 0o644)
		if err != nil {
			slog.Error("failed to open pyproject.toml", "err", err)
			os.Exit(1)
		}
		defer file.Close()

		if _, err := file.WriteString(pyprojectUvConfig); err != nil {
			slog.Error("failed to append uv config to pyproject.toml", "err", err)
			os.Exit(1)
		}
	},
}

func runPythonInitCommand(name string, args ...string) {
	stdout, err := cli_base.ExecCommand(name, args...)
	if err != nil {
		slog.Error("failed to run command", "command", name, "args", args, "err", err)
		os.Exit(1)
	}
	if stdout != "" {
		slog.Info("command completed", "command", name, "args", args, "stdout", stdout)
	}
}

func init() {
	pythonCmd.AddCommand(pythonInitCmd)
	rootCmd.AddCommand(pythonCmd)
}
