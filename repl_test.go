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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "charmander\ncharmander\ncharmander",
			expected: []string{"charmander", "charmander", "charmander"},
		},
		{
			input:    "\n\n   \n \n\n\n\nBulbasauR\tMewTWO\n\t\t\nraTTata     \n\n    ",
			expected: []string{"bulbasaur", "mewtwo", "rattata"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual output does not match length of expected output")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Actual word does not match expected word")
			}
		}
	}
}
