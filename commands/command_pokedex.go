package commands

import (
	"errors"
	"fmt"

	"github.com/FilipLusnia/gocli/config"
)

func callbackPokedex(cfg *config.CliConfig, args ...string) error {
	if len(cfg.CaughtPokemon) < 1 {
		return errors.New("you haven't caught any pokemon yet")
	}

	fmt.Printf("Pokemon in pokedex:\n")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
