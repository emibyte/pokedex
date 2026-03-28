package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  heLlo  worLD   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU   ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "gardeVoir  dragonite",
			expected: []string{"gardevoir", "dragonite"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, cse := range cases {
		actual := cleanInput(cse.input)

		if len(actual) != len(cse.expected) {
			t.Errorf("lengths to not match: len(actual): %d, len(expected): %d", len(actual), len(cse.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := cse.expected[i]

			if word != expectedWord {
				t.Errorf("actual: %s is not equal to expected: %s", word, expectedWord)
			}
		}
	}
}
