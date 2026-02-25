package main

import (
	"encoding/json"
	"os"

	"github.com/tbone317/pokedexcli/internal/pokeapi"
)

func clearCaughtPokemon(cfg *config, args ...string) error {
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
	saveCaughtPokemon(cfg)
	return nil
}

func saveCaughtPokemon(cfg *config, args ...string) error {
	f, err := os.Create("caught_pokemon.json")
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(cfg.caughtPokemon)
	if err != nil {
		return err
	}
	return nil
}

func loadCaughtPokemon(cfg *config) error {
	f, err := os.Open("caught_pokemon.json")
	if err != nil {
		if os.IsNotExist(err) {
			cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
			return nil // No file, no caught Pokemon
		}
		return err
	}
	defer f.Close()

	// Clear existing caught Pokemon
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg.caughtPokemon)
	if err != nil {
		return err
	}
	return nil
}
