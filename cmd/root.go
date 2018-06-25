package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Test Command!",
	Short: "Testing Cobra!",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cobra from cobratut is running!")
	},
}
