package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(" >")

	scanner.Scan()
	text := scanner.Text()

	fmt.Println("echoing input: ", text)
}
