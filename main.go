package main

import (
	"fly-pokeapi/src/cmd"
	"fmt"
	"os"
)

func main() {
	rootCmd := cmd.RootCommand()

	if len(os.Args) > 0 {
		commandName := os.Args[1]
		if commandName == "pokemon" {
			if len(os.Args) == 2 {
				fmt.Println("Please enter pokemon name")
			}

			if len(os.Args) == 3 {
				pokemonName := os.Args[1]
				fmt.Println(pokemonName)
				//create http request and ets
			}
		}
		if commandName == "attack" {
			//run command with function
		}

		if commandName == "stats" {
			//export data
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
