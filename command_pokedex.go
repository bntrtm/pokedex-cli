package main

import (
	"fmt"
)

func commandPokedex(c *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex {
		fmt.Println(fmt.Sprintf(" - %s", pokemon.Name))
	}
	return nil
}
