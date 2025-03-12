package commands

import "fmt"

type HelpCommand struct{}

func (h HelpCommand) Name() string {
	return "help"
}
func (h HelpCommand) Description() string {
	return "Displays a help message"
}
func (h HelpCommand) Execute() error {
	commands := GetCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.Name(), cmd.Description())
	}
	return nil
}
