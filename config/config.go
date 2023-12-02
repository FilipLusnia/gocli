package config

import (
	"time"

	"github.com/FilipLusnia/gocli/internal/pokeapi"
)

type CliConfig struct {
	Client              pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
}

var cfg = CliConfig{
	Client:              pokeapi.NewClient(time.Hour),
	NextLocationAreaURL: nil,
	PrevLocationAreaURL: nil,
}

func GetConfig() *CliConfig {
	return &cfg
}
