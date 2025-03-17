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

//func TestMapCommand(t *testing.T) {
//	testMap := []string{"canalave-city-area", "eterna-city-area", "pastoria-city-area", "sunyshore-city-area",
//		"sinnoh-pokemon-league-area", "oreburgh-mine-1f", "oreburgh-mine-b1f", "valley-windworks-area",
//		"eterna-forest-area", "fuego-ironworks-area", "mt-coronet-1f-route-207", "mt-coronet-2f", "mt-coronet-3f",
//		"mt-coronet-exterior-snowfall", "mt-coronet-exterior-blizzard", "mt-coronet-4f", "mt-coronet-4f-small-room", "mt-coronet-5f",
//		"mt-coronet-6f", "mt-coronet-1f-from-exterior"}
//
//	cases := []struct {
//		input    string
//		expected []string
//	}{
//		{
//			input:    "map",
//			expected: testMap,
//		},
//	}
//
//	for _, c := range cases {
//
//	}
//}
