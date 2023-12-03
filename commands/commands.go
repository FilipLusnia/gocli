package commands

import "github.com/FilipLusnia/gocli/config"

type cliCommand struct {
	name        string
	description string
	Callback    func(*config.CliConfig, ...string) error
}

type commandList map[string]*cliCommand

func GetCommands() commandList {
	return commandList{
		"map": {
			name:        "map",
			description: "Shows list of next location areas",
			Callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows list of previous location areas",
			Callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Shows list of pokemon in given location area",
			Callback:    callbackExplore,
		},
		"help": {
			name:        "help",
			description: "Shows available commands",
			Callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the CLI",
			Callback:    callbackExit,
		},
	}
}
