package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "flypoke",
	Short: "A basic CLI example",
	Long:  "A basic CLI example using Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Welcome to FlyPoke!!")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
