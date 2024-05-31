package cli

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var MigrateCommand cli.Command = cli.Command{
	Name:  "migrate",
	Usage: "Run database migrations",
	Action: func(c *cli.Context) error {
		// Menjalankan migrasi basis data
		cmd := exec.Command("your_migration_command", "--flags")
		cmd.Stdout = nil // Ganti dengan output stdout yang sesuai jika ada
		cmd.Stderr = nil // Ganti dengan output stderr yang sesuai jika ada
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run database migrations: %v", err)
		}

		log.Println("Database migrations completed")
		return nil
	},
}
