package main

import (
        "time"
        "github.com/bntrtm/pokedex-cli/internal/pokeapi"
        //"github.com/bntrtm/pokedex-cli/internal/pokecache"
)

func main() {
        //REPL time!
        client := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
        cfg := &config{
                pokedex: map[string]pokeapi.PokemonStat{},
                pClient: client,
        }
        startRepl(cfg)
}