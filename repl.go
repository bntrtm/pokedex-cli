package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/bntrtm/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type locationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type pokePage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
		pClient			pokeapi.Client
        urlNext         *string
        urlPrevious     *string
}

func cmdRegistry() map[string]cliCommand {
        cmdMap := map[string]cliCommand{
                "exit": {
                        name:        "exit",
                        description: "Exit the Pokedex",
                        callback:    commandExit,
                },
                "help": {
                        name:        "help",
                        description: "Show usage instructions",
                        callback:    commandHelp,
                },
                "map": {
                        name:         "map",
                        description:  "Display next 20 locations",
                        callback:     commandMap,
                },
                "mapb": {
                        name:         "mapb",
                        description:  "Display previous 20 locations",
                        callback:     commandMapb,
                },
				"explore": {
                        name:         "explore",
                        description:  "Given a location, list pokemon present",
                        callback:     commandExit,
                },
        }
        return cmdMap
}

func startRepl(cfg *config) {
	cmdMap := cmdRegistry()
	scanner := bufio.NewScanner(os.Stdin)
	for {
			fmt.Print("Pokedex > ")
			if scanner.Scan() {
					input := scanner.Text()
					userCmd := cleanInput(input)[0]
					if val, ok := cmdMap[userCmd]; ok {
							val.callback(cfg)
					} else {
							fmt.Println("Unknown command")
					}

			}
	}
}

func cleanInput(text string) []string {    
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}