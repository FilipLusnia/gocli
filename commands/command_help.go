package commands

import "fmt"

func callbackHelp() error {
	fmt.Println("")
	fmt.Println("Here is a list of available commands:")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - %s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println("")

	return nil
}
