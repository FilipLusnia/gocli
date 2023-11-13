package commands

import (
	"errors"
	"fmt"

	"github.com/FilipLusnia/gocli/config"
)

func callbackMap(cfg *config.CliConfig) error {
	if cfg.NextLocationAreaURL == nil {
		return errors.New("there is no next location")
	}

	nextLocation := config.GetConfig().NextLocationAreaURL
	resp, err := cfg.Client.ListLocationArea(nextLocation)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}

	fmt.Println("Locations:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapBack(cfg *config.CliConfig) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("there is no previous location")
	}

	prevLocation := config.GetConfig().PrevLocationAreaURL
	resp, err := cfg.Client.ListLocationArea(prevLocation)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}

	fmt.Println("Locations:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}
