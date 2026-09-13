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
			err := cmd.callback()
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	lowerText = strings.Trim(lowerText, " ")
	splitString := strings.Split(lowerText, " ")
	return splitString
}
