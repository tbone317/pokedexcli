package main

import (
	"fmt"
	"time"

	"github.com/tbone317/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	err := loadCaughtPokemon(cfg)
	if err != nil {
		fmt.Println("Error loading caught Pokemon:", err)
	}
	startRepl(cfg)
}
