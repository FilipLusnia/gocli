package commands

import (
	"fmt"

	"github.com/FilipLusnia/gocli/config"
)

func callbackHelp(cfg *config.CliConfig) error {
	fmt.Println("")
	fmt.Println("Here is a list of available commands:")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
