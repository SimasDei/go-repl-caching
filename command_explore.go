package main

import "errors"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Expected exactly one argument, location area name")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}

	println("Location area:")
	println(resp.Name)

	println("Encounters:")
	for _, encounter := range resp.PokemonEncounters {
		println(encounter.Pokemon.Name)
	}

	return nil
}
