package main

import (
	"errors"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		return err
	}

	println("Location areas:")

	for _, locationArea := range resp.Results {
		println(locationArea.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	return nil
}

func callbackMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("No previous location area URL")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)

	if err != nil {
		return err
	}

	println("Location areas:")

	for _, locationArea := range resp.Results {
		println(locationArea.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	return nil
}
