package config

import (
	"time"

	"github.com/FilipLusnia/gocli/internal/pokeapi"
)

type CliConfig struct {
	Client              pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
	CaughtPokemon       map[string]pokeapi.Pokemon
}

var cfg = CliConfig{
	Client:              pokeapi.NewClient(time.Hour),
	NextLocationAreaURL: nil,
	PrevLocationAreaURL: nil,
	CaughtPokemon:       make(map[string]pokeapi.Pokemon),
}

func GetConfig() *CliConfig {
	return &cfg
}
