package commands

import (
	"fmt"
	"github.com/maevlava/pokedex/model"
	"os"
)

type ExitCommand struct{}

func (e ExitCommand) Name() string {
	return "exit"
}

func (e ExitCommand) Description() string {
	return "Exit the Pokedex"
}

func (e ExitCommand) Execute(user *model.User, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
