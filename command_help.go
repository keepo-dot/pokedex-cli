package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range getCommands() {
		usageStr := fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
		fmt.Print(usageStr)
	}
	return nil
}
