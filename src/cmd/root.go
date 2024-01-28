package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flypoke",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Welcome to FlyPoke!!")
		},
	}

	// Register your commands here
	// cmd.AddCommand(ExampleCommand)

	return cmd
}
