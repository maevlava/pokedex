package tests

import (
	"fmt"
	command "github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPokedex_PokedexIsEmpty(t *testing.T) {
	user := &model.User{}
	cmd := &command.PokedexCommand{}
	std, r, w := CaptureStdOutput()
	err := cmd.Execute(user)
	cliOutput := RestoreStdOutput(std, r, w)

	assert.Errorf(t, err, "Empty Pokedex")
	assert.Empty(t, user.Pokedex, "It is not empty")
	assert.Contains(t, cliOutput.String(), "Your Pokedex is still empty")
}

func TestPokedex_PokedexShowsCapturedPokemon(t *testing.T) {
	expectedOutput := "Your Pokedex:\n\t- pidgey\n\t- caterpie\n"

	user := &model.User{}
	cmd := &command.PokedexCommand{}
	catchCmd := &command.CatchCommand{}

	catchCmd.Execute(user, "pidgey")
	catchCmd.Execute(user, "caterpie")

	std, r, w := CaptureStdOutput()
	err := cmd.Execute(user)
	if err != nil {
		fmt.Println(err)
	}
	cliOutput := RestoreStdOutput(std, r, w)

	assert.Contains(t, cliOutput.String(), expectedOutput)
}
