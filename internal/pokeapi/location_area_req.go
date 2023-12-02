package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}
	locationAreasResp := LocationAreasResp{}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Printf("cache hit with url %s", fullURL)
		err := json.Unmarshal(cacheData, &locationAreasResp)
		if err != nil {
			return locationAreasResp, err
		}

		return locationAreasResp, nil
	}
	fmt.Printf("cache miss with url %s", fullURL)

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
