package commands

type CliCommand interface {
	Name() string
	Description() string
	Execute() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": ExitCommand{},
		"help": HelpCommand{},
	}
}
