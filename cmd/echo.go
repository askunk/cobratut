package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for choing anything back.
	Ech echo's`,

	Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
	fmt.Println(strings.Join(args, " "))
}

func init() {
	RootCmd.AddCommand(cmdEcho)
}
