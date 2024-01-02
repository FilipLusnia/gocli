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
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch pokemon and add it to your pokedex",
			Callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Shows information about caught pokemon",
			Callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows list of caught pokemon",
			Callback:    callbackPokedex,
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
