package main

import "github.com/keepo-dot/pokedex-cli/internal/pokeapi"

type config struct {
	commandList   map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextPage      *string
	previousPage  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
