package commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type commandList map[string]*cliCommand

func GetCommands() commandList {
	return commandList{
		"help": {
			Name:        "help",
			Description: "Shows available commands",
			Callback:    callbackHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the CLI",
			Callback:    callbackExit,
		},
	}
}
