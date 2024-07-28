package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(" >")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

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

		err := command.callback(cfg, args...)

		if err != nil {
			fmt.Println("An error occurred: ", err.Error())
		}
	}

}

type CliCommand struct {
	Name        string
	Description string
	callback    func(*config, ...string) error
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
		"map-back": {
			Name:        "map-back",
			Description: "Lists the previous locations",
			callback:    callbackMapBack,
		},
		"explore": {
			Name:        "explore {location area name}",
			Description: "Explores a location",
			callback:    callbackExplore,
		},
	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}
