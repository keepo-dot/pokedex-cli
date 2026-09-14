package main

import (
	"time"

	"github.com/keepo-dot/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		commandList:   getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
