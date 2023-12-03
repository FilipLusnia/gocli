package commands

import (
	"errors"
	"fmt"

	"github.com/FilipLusnia/gocli/config"
)

func callbackMap(cfg *config.CliConfig, args ...string) error {
	resp, err := cfg.Client.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}
	if resp.Next == nil {
		return errors.New("there is no next location")
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapBack(cfg *config.CliConfig, args ...string) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("there is no previous location")
	}

	resp, err := cfg.Client.ListLocationAreas(cfg.PrevLocationAreaURL)
	if err != nil {
		return fmt.Errorf("error fetching areas: %w", err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}
