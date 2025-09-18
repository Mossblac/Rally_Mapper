package commands

var CommandMap map[string]CliCommand

type CliCommand struct {
	Name        string
	Description string
	Callback    func(string) error
}

func init() {
	CommandMap = map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "display help info",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Course Creator",
			Callback:    CommandExit,
		},
	}
}
