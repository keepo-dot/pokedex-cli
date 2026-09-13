package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	locationURL := "https://pokeapi.co/api/v2/location-area"

	if config.nextPage != "" {
		locationURL = config.nextPage
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
	if config.nextPage == "" {
		fmt.Println("You're on the last page.")
	}

	return nil
}
