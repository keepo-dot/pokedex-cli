package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.pokecache.Get(url); ok {
		cacheHit := LocationAreas{}
		err := json.Unmarshal(val, &cacheHit)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("error unmarshalling cache hit: %w", err)
		}
		return cacheHit, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error during http request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error during http response: %w", err)
	}
	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
	}
	c.pokecache.Add(url, jsonData)
	locationsResp := LocationAreas{}
	err = json.Unmarshal(jsonData, &locationsResp)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error during JSON Unmarshal: %w", err)
	}

	return locationsResp, nil
}
