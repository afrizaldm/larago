package main

import (
	"log"
	"os"

	c "simple-api/cmd/commands"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "artisan",
		Usage: "A simple CLI for building and serving the Go application",
		Commands: c.UseCommands(
			&c.BuildCommand,
			&c.ServeCommand,
			&c.BuildCleanCommand,
			&c.ConsoleClearCommand,
		),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
