package main

import (
	"bufio"
	"fmt"
	command "github.com/maevlava/pokedex/commands"
	util "github.com/maevlava/pokedex/utils"
	"os"
)

func main() {
	commands := command.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for {
		fmt.Printf("%s", "Pokedex > ")
		scanner.Scan()
		input = util.CleanInput(scanner.Text())
		cmd, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.Execute(input[1:]...); err != nil {
			fmt.Println(err)
		}
	}
}
