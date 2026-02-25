package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	saveCaughtPokemon(cfg)
	fmt.Println("Saving caught Pokemon... Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // unreachable, but required
}
