package cli

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var SeedCommand cli.Command = cli.Command{
	Name:  "seed",
	Usage: "Seed the database with initial or dummy data",
	Action: func(c *cli.Context) error {
		// Menjalankan seeding basis data
		cmd := exec.Command("your_seed_command", "--flags")
		cmd.Stdout = nil // Ganti dengan output stdout yang sesuai jika ada
		cmd.Stderr = nil // Ganti dengan output stderr yang sesuai jika ada
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}

		log.Println("Database seeding completed")
		return nil
	},
}
