package tests

import (
	"bytes"
	command "github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/model"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestCatch_CatchPokemon(t *testing.T) {
	captureOutput := func() (std *os.File, read *os.File, write *os.File) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		return oldStdout, r, w
	}
	restoreOutput := func(std *os.File, read *os.File, write *os.File) bytes.Buffer {
		write.Close()
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(read)
		os.Stdout = std
		return buf
	}

	cases := []struct {
		input          string
		expectedAlways string
		expectedEither []string
	}{
		{
			input:          "pikachu",
			expectedAlways: "throwing a Pokeball at pikachu...",
			expectedEither: []string{"pikachu escaped!", "pikachu was caught!"},
		},
		{
			input:          "squirtle",
			expectedAlways: "throwing a Pokeball at squirtle...",
			expectedEither: []string{"squirtle escaped!", "squirtle was caught!"},
		},
	}

	for _, c := range cases {
		std, r, w := captureOutput()

		cmd := &command.CatchCommand{}
		testUser := &model.User{}
		cmd.Execute(testUser, c.input)

		cliOutput := restoreOutput(std, r, w)
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

// TODO
//func TestCatch_PokemonAddedToUserPokdex(t *testing.T) {
//
//}
//func TestCatch_PokemonHasChanceToGetCatched(t *testing.T) {
//
//}
