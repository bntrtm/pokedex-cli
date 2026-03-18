package main

import (
	"fmt"
)

func commandExplore(c *config, args []string) error {
	var locName string
	if len(args) > 0 {
		locName = args[0]
	}

	locA, err := c.pClient.GetLocation(locName)
	if err != nil {
		return err
	}
	for _, enc := range locA.PokemonEncounters {
		pokeName := enc.Pokemon.Name
		fmt.Println(pokeName)
	}
	return nil
}
