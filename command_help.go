package main

import (
	"fmt"
	"strings"
)

func commandHelp(c *config, args []string) error {
	var helpString strings.Builder
	helpString.WriteString(`
Welcome to the Pokedex!
Usage:

`)
	for _, val := range cmdRegistry() {
		helpString.WriteString(val.name + ": " + val.description + "\n")
	}
	fmt.Println(helpString.String())
	return nil
}
