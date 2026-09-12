package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		output := fmt.Sprintf("Your command was: %s\n", inputCleaned[0])
		fmt.Print(output)
	}
}
