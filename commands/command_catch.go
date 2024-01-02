package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/FilipLusnia/gocli/config"
)

func callbackCatch(cfg *config.CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.Client.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error catching pokemon: %w", err)
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		fmt.Printf("failed to catch %s\n", pokemonName)
		return nil
	}
	cfg.CaughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
