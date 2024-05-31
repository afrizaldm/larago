package cli

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var BuildCommand cli.Command = cli.Command{
	Name:  "build",
	Usage: "Build the Go application",
	Action: func(c *cli.Context) error {
		// Hapus folder build jika ada
		if err := os.RemoveAll("build"); err != nil {
			log.Fatalf("Failed to remove build directory: %v", err)
		}

		// Membuat folder build
		if err := os.Mkdir("build", 0755); err != nil {
			log.Fatalf("Failed to create build directory: %v", err)
		}

		// Hapus folder build\database jika ada
		if err := os.RemoveAll("build\\database"); err != nil {
			log.Fatalf("Failed to remove build directory: %v", err)
		}

		// Membuat folder build\database
		if err := os.Mkdir("build\\database", 0755); err != nil {
			log.Fatalf("Failed to create build directory: %v", err)
		}

		// Menyalin folder public
		cmd := exec.Command("xcopy", "public", "build\\public", "/E", "/I", "/Y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to copy public directory: %v", err)
		}

		// Menyalin basis data dari database\database.sqlite ke build\database.sqlite
		cmd = exec.Command("xcopy", "database\\database.sqlite", "build\\database", "/Y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Gagal menyalin basis data: %v", err)
		}

		// Menyalin folder views
		cmd = exec.Command("xcopy", "resources\\views", "build\\views", "/E", "/I", "/Y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to copy views directory: %v", err)
		}

		// Membuat executable
		cmd = exec.Command("go", "build", "-o", "build\\main.exe", "./cmd/main.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to build application: %v", err)
		}

		log.Println("Build successful")
		return nil
	},
}
