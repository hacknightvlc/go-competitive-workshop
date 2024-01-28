package main

import (
	"fly-pokeapi/src/cmd"
	"fmt"
	"os"
)

func main() {
	rootCmd := cmd.RootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
