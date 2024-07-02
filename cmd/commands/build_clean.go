package commands

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var BuildCleanCommand cli.Command = cli.Command{
	Name:  "build:clean",
	Usage: "Clean the build directory",
	Action: func(c *cli.Context) error {
		// Hapus folder build jika ada
		if err := os.RemoveAll("build"); err != nil {
			log.Fatalf("Failed to remove build directory: %v", err)
		}

		// Tampilkan pesan sukses
		log.Println("Build directory cleaned")
		return nil
	},
}
