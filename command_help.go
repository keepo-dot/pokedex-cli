package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")
	fmt.Println()
	for _, cmd := range getCommands() {
		usageStr := fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
		fmt.Print(usageStr)
		fmt.Println()
	}
	return nil
}
