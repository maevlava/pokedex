package commands

type CliCommand interface {
	Name() string
	Description() string
	Execute() error
}

func GetCommands() map[string]CliCommand {
	pokeMapCommand := LoadMap()
	return map[string]CliCommand{
		"exit": ExitCommand{},
		"help": HelpCommand{},
		"map":  pokeMapCommand,
		"mapb": &PokeMapBackwardCommand{pm: pokeMapCommand},
	}
}
