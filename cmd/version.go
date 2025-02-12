package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DEFINE VERSION
const version = "0.1.0-beta"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of mdout",
	Long:  `Displays the current version of mdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mdout version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
