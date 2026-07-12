package player

import (
	"fmt"
	"testing"
)

func TestIntPow(t *testing.T) {
	cases := []struct {
		input    int
		power    int
		expected int
	}{
		{
			input:    1,
			power:    1,
			expected: 1,
		},
		{
			input:    2,
			power:    3,
			expected: 8,
		},
		{
			input:    0,
			power:    55,
			expected: 0,
		},
		{
			input:    98,
			power:    3,
			expected: 941192,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			result := intPow(c.input, c.power)
			success := result == c.expected
			if !success {
				t.Errorf("expected result to be %v, but was %v", c.expected, result)
			}
		})
	}
}
