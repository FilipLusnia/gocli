package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FilipLusnia/gocli/commands"
	"github.com/FilipLusnia/gocli/config"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())
		if len(cleanedText) == 0 {
			continue
		}

		availableCommands := commands.GetCommands()
		commandName := cleanedText[0]
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command - type \"help\" for available commands")
			continue
		}
		err := command.Callback(config.GetConfig())
		if err != nil {
			fmt.Printf("Error: %s", err)
			continue
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
