package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Expected exactly one argument, Pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	const threshold int = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("Pokemon %s got away", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon

	fmt.Println("You caught a Pokemon!")
	fmt.Printf("%v is now in your party\n", pokemon.Name)

	return nil
}
