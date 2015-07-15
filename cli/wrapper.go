package cli

import (
	"github.com/codegangsta/cli"
)

func Manage(c *cli.Context) {
	manage(c)
}

func GetCommands() []cli.Command {
	return commands
}

func RunWithCommand(newCommands []cli.Command) {
	commands = newCommands
	Run()
}
