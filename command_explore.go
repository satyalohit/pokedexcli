package main

import (
	"errors"
	"fmt"

)

func commandExplore (cfg *config, args ...string) error{
	
	if len(args) != 1{
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	locationArea,err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("pokemon in %s:\n", locationArea.Name)

       for _,pokemon :=range locationArea.PokemonEncounters {
		fmt.Printf("-%s\n",pokemon.Pokemon.Name)
	   }

	   return nil
}

