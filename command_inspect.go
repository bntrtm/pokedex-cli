package main

import (
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Enter a pokemon name to inspect!")
		return nil
	}

	target := args[0]

	if _, ok := c.pokedex[target]; !ok {
		fmt.Println("You have not caught that pokemon!")
		return nil
	}

	pokemon, err := c.pClient.GetPokemon(target)
	if err != nil {
		return err
	}
	pokemon.PrintStats()

	return nil
}
