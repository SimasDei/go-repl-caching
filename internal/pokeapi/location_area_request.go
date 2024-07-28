package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	val, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(val, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		println("Cache Hit: ", fullUrl)

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}
	c.cache.Add(fullUrl, data)
	println("Added to cache: ", fullUrl)

	return locationAreasResp, nil
}
