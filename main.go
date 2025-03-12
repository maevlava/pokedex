package main

import (
	"bufio"
	"fmt"
	util "github.com/maevlava/pokedex/utils"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for {
		fmt.Printf("%s", "Pokedex > ")
		scanner.Scan()
		input = util.CleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", input[0])
	}
}
