package main

import (
	"fmt"
)

func commandExplore(config *config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("location name must be included. ex: 'explore example'\n")
	}
	fmt.Printf("Exploring %s...\n", name)
	expRes, err := config.pokeapiClient.GetPokemon(name[0])
	if err != nil {
		return fmt.Errorf("error getting pokemon: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range expRes.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
