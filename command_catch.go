package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No pokemon specified. Enter a pokemon to throw at!")
		return nil
	}

	target := args[0]

	pokemon, err := c.pClient.GetPokemon(target)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	throwResult := rand.Intn(100)
	caught := false

	switch {
	case throwResult >= 95:
		caught = true
	case inBracket(throwResult, 80, 95):
		if pokemon.BaseExperience < 275 {
			caught = true
		}
	case inBracket(throwResult, 70, 80):
		if pokemon.BaseExperience < 220 {
			caught = true
		}
	case inBracket(throwResult, 50, 70):
		if pokemon.BaseExperience < 120 {
			caught = true
		}
	case inBracket(throwResult, 30, 50):
		if pokemon.BaseExperience < 90 {
			caught = true
		}
	case inBracket(throwResult, 0, 30):
		if pokemon.BaseExperience < 45 {
			caught = true
		}
	}

	if caught {
		c.pokedex[target] = pokemon
		fmt.Printf("You caught %s!\n", target)
	} else {
		fmt.Println("You missed the pokemon!")
	}

	return nil
}

func inBracket(input, moreOrEqualTo, LessThan int) bool {
	return input >= moreOrEqualTo && input < LessThan
}
