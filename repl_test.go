package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "",
			expected: []string{}},
		{input: " hello world ",
			expected: []string{"hello", "world"}},
		{input: "   multiple   spaces   ",
			expected: []string{"multiple", "spaces"}},
		{input: "UPPERCASE and lowercase",
			expected: []string{"uppercase", "and", "lowercase"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check if the length of the actual output matches the expected output
		// if the lengths don't match, use t.Errorf to report the failure
		if len(actual) != len(c.expected) {
			t.Fatalf("input %q: expected length %d but got %d", c.input, len(c.expected), len(actual))
		}
		for i := range c.expected {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("input %q (word %d): expected %q but got %q", c.input, i, expectedWord, word)
			}
		}
	}
}
