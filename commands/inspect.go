package commands

import (
	"errors"
	"fmt"
	"github.com/maevlava/pokedex/model"
)

type InspectCommand struct {
}

func (i *InspectCommand) Name() string {
	return "Inspect Command"
}

func (i *InspectCommand) Description() string {
	return "Inspect a pokemon status"
}

func (i *InspectCommand) Execute(user *model.User, args ...string) error {
	// It needs an argument
	if len(args) < 1 {
		return errors.New("Need 1 argument")
	}

	// get the pokemon
	inspectedPokemon := args[0]
	pokemon, err := getPokemon(user, inspectedPokemon)
	if err != nil {
		return err
	}

	// retrieve pokemon details
	inspectPokemon(pokemon)

	return nil
}

func getPokemon(user *model.User, inspectedPokemon string) (model.Pokemon, error) {
	var targetPokemon model.Pokemon

	// Read if the user has a pokemon
	if len(user.Pokedex) == 0 {
		fmt.Printf("You have not caught any pokemon yet.")
		return model.Pokemon{}, errors.New("No pokemon found")
	}

	// Read if the user has the pokemon
	found := false
	for _, pokemon := range user.Pokedex {
		if pokemon.Name == inspectedPokemon {
			found = true
			targetPokemon = pokemon
			break
		}
	}

	// If not found, return message
	if !found {
		fmt.Printf("You have not caught that pokemon")
	}
	return targetPokemon, nil
}
func inspectPokemon(pokemon model.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("\t -%s:%v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t - %s\n", t.Type.Name)
	}
}
