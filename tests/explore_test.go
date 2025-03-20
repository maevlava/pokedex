package tests

import (
	"fmt"
	command "github.com/maevlava/pokedex/commands"
	"github.com/maevlava/pokedex/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplore_MustHaveArgs(t *testing.T) {
	cmd := &command.ExploreCommand{}
	err := cmd.Execute(&model.User{})
	assert.Error(t, err, "Expected an error when no arguments are provided")
}
func TestExplore_ListPokemonInArea(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "canalave-city-area",
			output: []string{"tentacool", "tentacruel"},
		},
		{
			input:  "eterna-city-area",
			output: []string{"psyduck", "golduck"},
		},
	}
	for _, c := range cases {
		cmd := &command.ExploreCommand{}

		cmd.Execute(&model.User{}, c.input)
		pokemons := []string{}
		for _, pokemon := range cmd.Pokemons {
			pokemons = append(pokemons, pokemon.Name)
		}

		actual := pokemons
		assert.Subset(t, actual, c.output, fmt.Sprintf("Pokemons are not equal: %s, %s", actual, c.output))
	}
}
