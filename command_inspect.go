package main

import (
	"fmt"
)

func commandInspect(c *config, args []string) error {
	target := "nothing"
	if len(args) > 0 {
		target = args[0]
	} else {
		fmt.Println("Enter a pokemon name to inspect!")
		return nil
	}

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
