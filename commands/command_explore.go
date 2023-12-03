package commands

import (
	"errors"
	"fmt"

	"github.com/FilipLusnia/gocli/config"
)

func callbackExplore(cfg *config.CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.Client.GetLocationArea(locationAreaName)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}

	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
