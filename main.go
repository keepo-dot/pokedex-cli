package main

import (
	"time"

	"github.com/keepo-dot/pokedex-cli/internal/pokeapi"
	"github.com/keepo-dot/pokedex-cli/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(5*time.Second, pokeCache)
	config := &config{
		commandList:   getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
