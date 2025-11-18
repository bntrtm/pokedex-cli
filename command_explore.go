package main

import (
	"fmt"
)

func commandExplore(c *config, args []string) error {
	locName := "1"
	if len(args) > 0 {
		locName = args[0]
	}

	locA, err := c.pClient.GetLocation(locName)
	if err != nil {
		return err
	}
	pokemon := []string{}
	for _, enc := range locA.PokemonEncounters {
		pokeName := enc.Pokemon.Name
		fmt.Println(pokeName)
		pokemon = append(pokemon, pokeName)
	}
	return nil
}
