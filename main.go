package main

import (
        "time"
        "github.com/bntrtm/pokedex-cli/internal/pokeapi"
        //"github.com/bntrtm/pokedex-cli/internal/pokecache"
)

func main() {
        //REPL time!
        client := pokeapi.NewClient(5 * time.Second)
        cfg := &config{
                pClient: client,
        }
        startRepl(cfg)
}