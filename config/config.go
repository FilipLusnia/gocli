package config

import "github.com/FilipLusnia/gocli/internal/pokeapi"

type CliConfig struct {
	Client              pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
}

func GetConfig() *CliConfig {
	cfg := CliConfig{
		Client:              pokeapi.NewClient(),
		NextLocationAreaURL: nil,
		PrevLocationAreaURL: nil,
	}

	return &cfg
}
