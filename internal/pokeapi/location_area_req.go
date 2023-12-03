package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}
	locationAreasResp := LocationAreasResp{}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		err := json.Unmarshal(cacheData, &locationAreasResp)
		if err != nil {
			return locationAreasResp, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreasResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreasResp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return locationAreasResp, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreasResp, err
	}

	err = json.Unmarshal(respData, &locationAreasResp)
	if err != nil {
		return locationAreasResp, err
	}
	c.cache.Add(fullURL, respData)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (SingleLocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint
	locationArea := SingleLocationArea{}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			return locationArea, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationArea, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return locationArea, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationArea, err
	}

	err = json.Unmarshal(respData, &locationArea)
	if err != nil {
		return locationArea, err
	}
	c.cache.Add(fullURL, respData)

	return locationArea, nil
}
