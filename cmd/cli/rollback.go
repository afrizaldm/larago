package cli

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var RollbackCommand cli.Command = cli.Command{
	Name:  "rollback",
	Usage: "Rollback the last database migration",
	Action: func(c *cli.Context) error {
		// Melakukan rollback migrasi basis data
		cmd := exec.Command("your_rollback_command", "--flags")
		cmd.Stdout = nil // Ganti dengan output stdout yang sesuai jika ada
		cmd.Stderr = nil // Ganti dengan output stderr yang sesuai jika ada
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to rollback database migration: %v", err)
		}

		log.Println("Database migration rollback completed")
		return nil
	},
}
