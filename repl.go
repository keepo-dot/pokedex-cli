package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	lowerText = strings.Trim(lowerText, " ")
	splitString := strings.Split(lowerText, " ")
	return splitString
}
