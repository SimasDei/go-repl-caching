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

		fmt.Println("echoing input: ", text)
	}

}

func cleanInput(input string) []string {
	// clean the input string
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}
