package commands

import (
	"errors"
	"fmt"
	"github.com/maevlava/pokedex/model"
)

type PokedexCommand struct{}

func (p *PokedexCommand) Name() string {
	return "pokedex"
}

func (p *PokedexCommand) Description() string {
	return "Display all captured pokemon"
}

func (p *PokedexCommand) Execute(user *model.User, args ...string) error {
	if len(user.Pokedex) == 0 {
		fmt.Println("Your Pokedex is still empty")
		return errors.New("Empty Pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range user.Pokedex {
		fmt.Printf("\t- %s\n", pokemon.Name) // Removed extra spaces
	}
	return nil
}
