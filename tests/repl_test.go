package tests

import (
	"github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/internal/pokecache"
	"github.com/maevlava/pokedex/model"
	util "github.com/maevlava/pokedex/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestMapCommand(t *testing.T) {
	testMapB := []string{"canalave-city-area", "eterna-city-area", "pastoria-city-area", "sunyshore-city-area",
		"sinnoh-pokemon-league-area", "oreburgh-mine-1f", "oreburgh-mine-b1f", "valley-windworks-area",
		"eterna-forest-area", "fuego-ironworks-area", "mt-coronet-1f-route-207", "mt-coronet-2f", "mt-coronet-3f",
		"mt-coronet-exterior-snowfall", "mt-coronet-exterior-blizzard", "mt-coronet-4f", "mt-coronet-4f-small-room", "mt-coronet-5f",
		"mt-coronet-6f", "mt-coronet-1f-from-exterior"}
	testMap := []string{
		"mt-coronet-1f-route-216",
		"mt-coronet-1f-route-211",
		"mt-coronet-b1f",
		"great-marsh-area-1",
		"great-marsh-area-2",
		"great-marsh-area-3",
		"great-marsh-area-4",
		"great-marsh-area-5",
		"great-marsh-area-6",
		"solaceon-ruins-2f",
		"solaceon-ruins-1f",
		"solaceon-ruins-b1f-a",
		"solaceon-ruins-b1f-b",
		"solaceon-ruins-b1f-c",
		"solaceon-ruins-b2f-a",
		"solaceon-ruins-b2f-b",
		"solaceon-ruins-b2f-c",
		"solaceon-ruins-b3f-a",
		"solaceon-ruins-b3f-b",
		"solaceon-ruins-b3f-c",
	}

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "mapb",
			expected: testMapB,
		},
		{
			input:    "map",
			expected: testMap,
		},
	}

	//test case 1
	t.Run(cases[0].input, func(t *testing.T) {
		cache := pokecache.NewCache(10 * time.Minute)
		cmd := commands.LoadMap(cache)
		err := cmd.Execute(&model.User{})
		assert.NoErrorf(t, err, "Execute() Should not return an error")
		err = cmd.Execute(&model.User{})
		assert.NoErrorf(t, err, "Execute() Should not return an error")

		cmdBackward := commands.PokeMapBackwardCommand{Pm: cmd}
		err = cmdBackward.Execute(&model.User{})
		assert.NoErrorf(t, err, "Execute() Should not return an error")

		var actualNames []string
		for _, loc := range cmd.PokeMaps {
			actualNames = append(actualNames, loc.Name)
		}

		assert.Subset(t, actualNames, testMapB, cases[0].expected, "Returned map names do not match expected values")
	})

}
