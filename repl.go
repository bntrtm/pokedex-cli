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