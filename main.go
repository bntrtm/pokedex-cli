package main

import (
        "strings"
        "io"
        "os"
        "bufio"
        "fmt"
        "net/http"
        "encoding/json"
        //"github.com/bntrtm/pokedex-cli/internal/pokecache"
)

func cleanInput(text string) []string {    
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
        urlNext         string
        urlPrevious     string
}

func (c *config) init() {
        c.urlNext = "https://pokeapi.co/api/v2/location-area/"
        c.urlPrevious = ""
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
        }
        return cmdMap
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

func commandMap(c *config) error {

        client := &http.Client{}
        req, err := http.NewRequest("GET", c.urlNext, nil)
        if err != nil {
                fmt.Println(err)
        }
        resp, err := client.Do(req)
        if err != nil {
                fmt.Println(err)
        }
        defer resp.Body.Close()

        if resp.StatusCode >= 300 {
                fmt.Println(fmt.Sprintf("HTTP error %d for request at /%s", resp.StatusCode, c.urlNext))
        }

        jsonData, err := io.ReadAll(resp.Body)
        if err != nil {
                fmt.Println(err)
        }

        var page pokePage
        err = json.Unmarshal(jsonData, &page)
        if err != nil {
                fmt.Println(err)
        }

        c.urlNext = page.Next
        verifyPrevious, ok := page.Previous.(string)
        if !ok && c.urlPrevious != "" {
                fmt.Println("'page.Previous' is not a string")
        } else {
                c.urlPrevious = verifyPrevious
        }

        for _, result := range page.Results {
                fmt.Println(result.Name)
        }

        return nil

}

func commandMapb(c *config) error {

        if c.urlPrevious == "" {
                fmt.Println("You're on the first page!")
                return nil
        }

        client := &http.Client{}
        req, err := http.NewRequest("GET", c.urlPrevious, nil)
        if err != nil {
                fmt.Println(err)
        }
        resp, err := client.Do(req)
        if err != nil {
                fmt.Println(err)
        }
        defer resp.Body.Close()

        if resp.StatusCode >= 300 {
                fmt.Println(fmt.Sprintf("HTTP error %d for request at /%s", resp.StatusCode, c.urlPrevious))
        }

        jsonData, err := io.ReadAll(resp.Body)
        if err != nil {
                fmt.Println(err)
        }

        var page pokePage
        err = json.Unmarshal(jsonData, &page)
        if err != nil {
                fmt.Println(err)
        }

        c.urlNext = page.Next
        if page.Previous == nil {
                c.urlPrevious = ""
        } else {
                verifyPrevious, ok := page.Previous.(string)
                if ok {
                        c.urlPrevious = verifyPrevious
                }
        }
        
        for _, result := range page.Results {
                fmt.Println(result.Name)
        }

        return nil
}

func commandExit(c *config) error {
        fmt.Println("Closing the Pokedex... Goodbye!")
        os.Exit(0)
        return nil
}

func commandHelp(c *config) error {
        helpString := `
Welcome to the Pokedex!
Usage:

`
        for _, val := range cmdRegistry() {
                helpString += val.name + ": " + val.description + "\n"
        }
        fmt.Println(helpString)
        return nil
}

func main() {
        //REPL time!
        cmdMap := cmdRegistry()
        cfg := config{}
        cfg.init()
        scanner := bufio.NewScanner(os.Stdin)
        for {
                fmt.Print("Pokedex > ")
                if scanner.Scan() {
                        input := scanner.Text()
                        userCmd := cleanInput(input)[0]
                        if val, ok := cmdMap[userCmd]; ok {
                                val.callback(&cfg)
                        } else {
                                fmt.Println("Unknown command")
                        }

                }
        }
}