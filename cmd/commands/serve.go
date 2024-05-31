package commands

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var ServeCommand cli.Command = cli.Command{
	Name:  "serve",
	Usage: "Serve the Go application",
	Action: func(c *cli.Context) error {
		// Menjalankan aplikasi dalam mode debug
		cmd := exec.Command("go", "run", "./cmd/main.go", "--debug")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to serve application: %v", err)
		}

		return nil
	},
}
