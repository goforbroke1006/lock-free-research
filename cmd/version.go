package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of lock-free-research",
		Long:  `TODO: write me`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1")
		},
	}

	rootCmd.AddCommand(versionCmd)
}
