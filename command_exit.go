package main

import "os"

func callbackExit(cfg *config, args ...string) error {
	println("Goodbye!")
	os.Exit(0)

	return nil
}
