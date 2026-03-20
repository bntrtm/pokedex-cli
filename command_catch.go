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
	case throwResult >= ThrowIncredible:
		caught = true
	case inBracket(throwResult, ThrowGreat, ThrowIncredible):
		if pokemon.BaseExperience < DifficultyVeryHard {
			caught = true
		}
	case inBracket(throwResult, ThrowGood, ThrowGreat):
		if pokemon.BaseExperience < DifficultyHard {
			caught = true
		}
	case inBracket(throwResult, ThrowOkay, ThrowGood):
		if pokemon.BaseExperience < DifficultyMedium {
			caught = true
		}
	case inBracket(throwResult, ThrowLazy, ThrowOkay):
		if pokemon.BaseExperience < DifficultyEasy {
			caught = true
		}
	case inBracket(throwResult, DifficultyNone, ThrowLazy):
		if pokemon.BaseExperience < DifficultyVeryEasy {
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
