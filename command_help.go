package main

import (
	"fmt"
)

func commandHelp(c *config, args []string) error {
	helpString := `
Welcome to the Pokedex!
Usage:

`
	for _, val := range cmdRegistry() {
		helpString += val.name + ": " + val.description + "\n"
	}
	fmt.Println(helpString)
	return nil
}
