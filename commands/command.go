package commands

import (
	"github.com/maevlava/pokedex/internal/pokecache"
	"time"
)

type CliCommand interface {
	Name() string
	Description() string
	Execute() error
}

func GetCommands() map[string]CliCommand {
	cache := pokecache.NewCache(10 * time.Minute)
	pokeMapCommand := LoadMap(cache)

	return map[string]CliCommand{
		"exit": ExitCommand{},
		"help": HelpCommand{},
		"map":  pokeMapCommand,
		"mapb": &PokeMapBackwardCommand{Pm: pokeMapCommand},
	}
}
