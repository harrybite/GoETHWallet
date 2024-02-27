package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "print the version of cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.2")
	},
}

func init() {
	rootCmd.AddCommand(version)
}
