package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct{}

func (e ExitCommand) Name() string {
	return "exit"
}

func (e ExitCommand) Description() string {
	return "Exit the Pokedex"
}

func (e ExitCommand) Execute() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
