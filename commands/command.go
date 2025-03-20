package commands

import (
	"github.com/maevlava/pokedex/internal/pokecache"
	"github.com/maevlava/pokedex/model"
	"time"
)

type CliCommand interface {
	Name() string
	Description() string
	Execute(user *model.User, args ...string) error
}

func GetCommands() map[string]CliCommand {
	cache := pokecache.NewCache(10 * time.Minute)
	pokeMapCommand := LoadMap(cache)

	return map[string]CliCommand{
		"exit":    ExitCommand{},
		"help":    HelpCommand{},
		"map":     pokeMapCommand,
		"mapb":    &PokeMapBackwardCommand{Pm: pokeMapCommand},
		"explore": &ExploreCommand{},
		"catch":   &CatchCommand{},
	}
}
