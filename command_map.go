package main

import "github.com/SimasDei/go-repl-caching/internal/pokeapi"

func callbackMap() error {
	pokeApiClient := pokeapi.NewClient()

	resp, err := pokeApiClient.ListLocationAreas()

	if err != nil {
		panic(err)
	}

	println("Location areas:")

	for _, locationArea := range resp.Results {
		println(locationArea.Name)
	}

	return nil
}
