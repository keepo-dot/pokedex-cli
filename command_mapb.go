package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(config *config) error {
	locationURL := "https://pokeapi.co/api/v2/location-area"

	if config.previousPage != "" {
		locationURL = config.previousPage
	}
	if config.previousPage == "" {
		fmt.Println("You're on the first page.")
		return nil
	}
	res, err := http.Get(locationURL)
	if err != nil {
		return fmt.Errorf("error during http request: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	locRes := locationResponse{}

	err = json.Unmarshal(body, &locRes)
	if err != nil {
		return fmt.Errorf("error during JSON Unmarshal: %w", err)
	}
	config.nextPage = locRes.Next
	config.previousPage = locRes.Previous

	for _, location := range locRes.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
