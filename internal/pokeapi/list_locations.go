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
	locationsResp := LocationAreas{}
	err = json.Unmarshal(jsonData, &locationsResp)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error during JSON Unmarshal: %w", err)
	}
	return locationsResp, nil
}
