package tests

import (
	util "github.com/maevlava/pokedex/utils"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"Hello", "World"},
		},
		{
			input:    "one two three",
			expected: []string{"One", "Two", "Three"},
		},
		{
			input:    "",
			expected: []string{""},
		},
	}
	for _, c := range cases {
		actual := util.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("CleanInput(%q): length mismatch: got %d, expected %d", c.input, len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("CleanInput(%q): element %d mismatch: got %q, expected %q", c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
