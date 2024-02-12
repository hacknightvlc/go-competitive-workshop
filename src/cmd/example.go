package cmd

import (
	"github.com/spf13/cobra"
)

var ExampleCmd = &cobra.Command{
	Use:   "example",
	Short: "Example command",
	Long:  "Example command longer description",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Example command")
	},
}

func init() {
	RootCmd.AddCommand(ExampleCmd)
}
