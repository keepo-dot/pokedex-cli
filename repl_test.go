package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Bulbasaur is a cutie!",
			expected: []string{"bulbasaur", "is", "a", "cutie!"},
		},
		{
			input:    "SNORLAX iS tHe BeSt PoKeMoN in the game     ",
			expected: []string{"snorlax", "is", "the", "best", "pokemon", "in", "the", "game"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of input does not match expected output.\n actual: %v \n expected: %v\n", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word does not match expected word. \n actual: %s \n expected: %s\n", word, expectedWord)
			}
		}
	}
}
