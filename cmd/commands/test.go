package commands

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var TestCommand cli.Command = cli.Command{
	Name:  "test",
	Usage: "Run unit tests for the application",
	Action: func(c *cli.Context) error {
		// Menjalankan tes unit
		cmd := exec.Command("go", "test", "./...")
		cmd.Stdout = nil // Ganti dengan output stdout yang sesuai jika ada
		cmd.Stderr = nil // Ganti dengan output stderr yang sesuai jika ada
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run unit tests: %v", err)
		}

		log.Println("Unit tests completed")
		return nil
	},
}
