package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon" + "/" + pokemonName
	fullUrl := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("response error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
