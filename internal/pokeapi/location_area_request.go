package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area?offset=0&limit=20"
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

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area" + "/" + locationAreaName
	fullUrl := baseUrl + endpoint

	val, ok := c.cache.Get(fullUrl)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(val, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		println("Cache Hit: ", fullUrl)

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("response error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullUrl, data)
	println("Added to cache: ", fullUrl)

	return locationArea, nil
}
