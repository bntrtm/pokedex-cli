package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	target := "nothing"
	if len(args) > 0 {
		target = args[0]
	} else {
		fmt.Println("No pokemon specified. Enter a pokemon to throw at!")
		return nil
	}

	pokemon, err := c.pClient.GetPokemon(target)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Throwing a Pokeball at %s...", target))

	throwResult := rand.Intn(100)
	caught := false

	switch {
	case throwResult >= 95:
		caught = true
	case throwResult >= 80 && throwResult < 95:
		if pokemon.BaseExperience < 275 {
			caught = true
		}
	case throwResult >= 70 && throwResult < 80:
		if pokemon.BaseExperience < 220 {
			caught = true
		}
	case throwResult >= 50 && throwResult < 70:
		if pokemon.BaseExperience < 120 {
			caught = true
		}
	case throwResult >= 30 && throwResult < 50:
		if pokemon.BaseExperience < 90 {
			caught = true
		}
	case throwResult >= 0 && throwResult < 30:
		if pokemon.BaseExperience < 45 {
			caught = true
		}
	}

	if caught {
		c.pokedex[target] = pokemon
		fmt.Println(fmt.Sprintf("You caught %s!", target))
	} else {
		fmt.Println("You missed the pokemon!")
	}

	return nil
}
