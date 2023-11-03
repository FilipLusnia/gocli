package commands

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

type commandList map[string]*cliCommand

func GetCommands() commandList {
	return commandList{
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
