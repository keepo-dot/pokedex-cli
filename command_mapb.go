package main

import (
	"fmt"
)

func commandMapB(config *config, name ...string) error {
	if config.previousPage == nil {
		fmt.Print("you're on the first page\n")
		return nil
	}

	locRes, err := config.pokeapiClient.GetLocations(config.previousPage)
	if err != nil {
		return fmt.Errorf("error getting locations: %w", err)
	}

	config.nextPage = locRes.Next
	config.previousPage = locRes.Previous

	for _, location := range locRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
