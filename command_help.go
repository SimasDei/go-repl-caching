package main

import "fmt"

func callbackHelp() error {
	availableCommands := getCommands()

	for _, command := range availableCommands {
		fmt.Printf("%s - %s\n", command.Name, command.Description)
	}

	return nil
}
