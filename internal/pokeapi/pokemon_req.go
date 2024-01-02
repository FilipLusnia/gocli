package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint
	pokemon := Pokemon{}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		err := json.Unmarshal(cacheData, &pokemon)
		if err != nil {
			return pokemon, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return pokemon, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return pokemon, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(respData, &pokemon)
	if err != nil {
		return pokemon, err
	}
	c.cache.Add(fullURL, respData)

	return pokemon, nil
}
