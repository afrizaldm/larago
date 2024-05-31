package commands

import "github.com/urfave/cli/v2"

func UseCommands(commands ...*cli.Command) []*cli.Command {
	return commands
}
