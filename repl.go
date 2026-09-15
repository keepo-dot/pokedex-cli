package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		if scanner.Err() != nil {
			err := fmt.Sprintf("scanner error: %s\n", scanner.Err().Error())
			fmt.Print(err)
		}
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		cmd, exists := getCommands()[inputCleaned[0]]
		if exists {
			err := cmd.callback(config, inputCleaned[1:]...)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	lowerText = strings.Trim(lowerText, " ")
	splitString := strings.Split(lowerText, " ")
	return splitString
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 areas in the Pokemon world. \nSubsequent calls show the next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas in the Pokemon world.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon in a given area.",
			callback:    commandExplore,
		},
	}
}
