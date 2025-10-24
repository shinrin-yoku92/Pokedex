package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/shinrin_yoku92/Pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catchChance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchChance > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		if cfg.caughtPokemon == nil {
			cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
		}
		cfg.caughtPokemon[pokemon.Name] = pokemon
	}
	return nil
}
