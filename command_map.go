package main

import (
	"fmt"
)

func commandMap(config *config, name ...string) error {
	locRes, err := config.pokeapiClient.GetLocations(config.nextPage)
	if err != nil {
		return err
	}

	config.nextPage = locRes.Next
	config.previousPage = locRes.Previous

	for _, location := range locRes.Results {
		fmt.Println(location.Name)
	}
	if config.nextPage == nil {
		fmt.Println("You're on the last page.")
	}
	return nil
}
