package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloName string

var helloCMD = &cobra.Command{
	Use:   "hello",
	Short: "Print a greetings name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "Hello there Mr. %s!\n", helloName)
	},
}

func init() {
	helloCMD.Flags().StringVarP(&helloName, "name", "n", "world", "who to greet")
	rootCmd.AddCommand(helloCMD)
}
