package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lock-free-research",
	Short: "TODO: write me",
	Long: `TODO: write me
TODO: write me
TODO: write me`,
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Request handlers",
	Long: `* HTTP handler
* Data bus command executor
* REST API handler`,
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Daemon programs",
	Long: `* Periodic job
* Data bus handlers
* Events listeners`,
}

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Additional soft",
	Long: `* CLIs
* Manual fixers`,
}

func Execute() error {
	rootCmd.AddCommand(apiCmd, daemonCmd, utilCmd)
	return rootCmd.Execute()
}
