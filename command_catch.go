package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50

	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println("Throwing a Pokeball at pikachu...")
	if randNum > threshold {

		return fmt.Errorf(" %s was escaped", pokemonName)

	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught", pokemonName)

	return nil
}
