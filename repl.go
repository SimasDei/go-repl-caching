package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(" >")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()

		fmt.Println("echoing input: ", cleaned)
	}

}

type CliCommand struct {
	Name        string
	Description string
	callback    func() error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Shows the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the program",
			callback:    callbackExit,
		},
		"map": {
			Name:        "map",
			Description: "Lists the locations",
			callback:    callbackMap,
		},
	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}
