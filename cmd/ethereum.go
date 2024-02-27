package cmd

import (
	// "fmt"

	"github.com/spf13/cobra"
)

var ethereum = &cobra.Command{
	Use:   "eth",
	Short: "extract ethereum wallet address form private key",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(ethereum)
}
