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
			err := cmd.callback(config)
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
