package commands

import (
	"os"

	"github.com/FilipLusnia/gocli/config"
)

func callbackExit(cfg *config.CliConfig, args ...string) error {
	os.Exit(0)
	return nil
}
