package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.caughtPokemon
	if len(pokedex) == 0 {
		fmt.Println("No caught Pokemon yet.")
		return nil
	}

	fmt.Println("Caught Pokemon:")
	for name := range pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
