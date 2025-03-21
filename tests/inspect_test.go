package tests

import (
	command "github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInspect_PokemonDetails(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input: "pidgey",
			output: []string{
				"Name: pidgey",
				"Height: 3",
				"Weight: 18",
				"Stats:",
				"\t -hp:40",
				"\t -attack:45",
				"\t -defense:40",
				"\t -special-attack:35",
				"\t -special-defense:35",
				"\t -speed:56",
				"Types:",
				"\t - normal",
				"\t - flying",
			},
		},
	}
	for _, c := range cases {
		user := &model.User{}
		cmd := &command.InspectCommand{}

		catchCmd := &command.CatchCommand{}
		catchCmd.Execute(user, c.input)

		std, r, w := CaptureStdOutput()
		cmd.Execute(user, c.input)
		cliOutput := RestoreStdOutput(std, r, w)

		for _, expectedLine := range c.output {
			assert.Contains(t, cliOutput.String(), expectedLine, "Expected output to contain: %s", expectedLine)
		}

	}
}

func TestInspect_PokemonDoesNotExist(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "pikachu",
			output: "You have not caught that pokemon",
		},
	}
	user := &model.User{}

	cmd := &command.InspectCommand{}
	catchCmd := &command.CatchCommand{}

	catchCmd.Execute(user, "squirtle")

	for _, testCase := range cases {
		std, r, w := CaptureStdOutput()
		cmd.Execute(user, testCase.input)
		cliOutput := RestoreStdOutput(std, r, w)
		assert.Equal(t, testCase.output, cliOutput.String())
	}
}

func TestInspect_EmptyPokedex(t *testing.T) {
	user := &model.User{}
	cmd := &command.InspectCommand{}

	std, r, w := CaptureStdOutput()
	cmd.Execute(user, "pikachu")
	cliOutput := RestoreStdOutput(std, r, w)

	assert.Equal(t, "You have not caught any pokemon yet.", cliOutput.String())
}
