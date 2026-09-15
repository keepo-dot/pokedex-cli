package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(areaName string) (ExploreResponse, error) {
	url := baseURL + "/location-area/" + areaName
	//if pageUrl != nil {
	//	url = *pageUrl
	//}
	if val, ok := c.pokecache.Get(url); ok {
		cacheHit := ExploreResponse{}
		err := json.Unmarshal(val, &cacheHit)
		if err != nil {
			return ExploreResponse{}, fmt.Errorf("error unmarshalling cache hit: %w", err)
		}
		return cacheHit, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error during http request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error during http response: %w", err)
	}
	if res.StatusCode > 299 {
		return ExploreResponse{}, fmt.Errorf("response failed with status code: %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error reading response body: %w", err)
	}
	c.pokecache.Add(url, jsonData)
	expResp := ExploreResponse{}
	err = json.Unmarshal(jsonData, &expResp)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error during json unmarshal: %w", err)
	}
	return expResp, nil
}
