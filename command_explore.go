package main

import (
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location specified to explore")
	}

	locName := args[0]

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
