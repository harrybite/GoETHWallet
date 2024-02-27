package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli toop for getting wallet address from ethereum private key",
	Long: `command line interface which extract the wallet address from the private key compatible for ethereum.
	this command only works for ethereum private key`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Hugo is a very fast static site generator2")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
