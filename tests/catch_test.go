package tests

import (
	command "github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCatch_CatchPokemon(t *testing.T) {
	cases := []struct {
		input          string
		expectedAlways string
		expectedEither []string
	}{
		{
			input:          "pikachu",
			expectedAlways: "Throwing a Pokeball at pikachu...",
			expectedEither: []string{"pikachu escaped!", "pikachu was caught!"},
		},
		{
			input:          "squirtle",
			expectedAlways: "Throwing a Pokeball at squirtle...",
			expectedEither: []string{"squirtle escaped!", "squirtle was caught!"},
		},
	}

	for _, c := range cases {
		std, r, w := CaptureStdOutput()

		cmd := &command.CatchCommand{}
		testUser := &model.User{}
		cmd.Execute(testUser, c.input)

		cliOutput := RestoreStdOutput(std, r, w)
		assert.Contains(t, cliOutput.String(), c.expectedAlways)

		found := false
		for _, possibleOutput := range c.expectedEither {
			if strings.Contains(cliOutput.String(), possibleOutput) {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected one of the outputs to be present")
	}

}

func TestCatch_PokemonAddedToUserPokdex(t *testing.T) {
	cmd := &command.CatchCommand{}
	testUser := &model.User{}
	cmd.Execute(testUser, "pikachu")
	assert.Contains(t, testUser.Pokedex[0].Name, "pikachu")
}
