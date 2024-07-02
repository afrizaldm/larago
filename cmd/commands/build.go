package commands

import (
	"io"
	"log"
	"os"
	"os/exec"
	"simple-api/config"
	"time"

	"github.com/urfave/cli/v2"
)

var BuildCommand cli.Command = cli.Command{
	Name:  "build",
	Usage: "Build the Go application",
	Action: func(c *cli.Context) error {

		log.Println("Building...")

		// Start timer
		start := time.Now()

		// clear
		Clean()
		SetSQlite()
		SetLogger()
		SetENV()
		SetStatic()
		SetView()
		Build()

		duration := time.Since(start)

		log.Println("Build Successful In", duration)
		return nil
	},
}

func Clean() {
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

	// Hapus folder build\logs jika ada
	if err := os.RemoveAll("build\\logs"); err != nil {
		log.Fatalf("Failed to remove build directory: %v", err)
	}
}

func SetENV() {
	// Buka file asli
	srcFile, err := os.Open(".\\.env")
	if err != nil {
		log.Fatalf("Gagal membuka file asli: %v", err)
	}
	defer srcFile.Close()

	// Buat file tujuan
	dstFile, err := os.Create("build\\.env")
	if err != nil {
		log.Fatalf("Gagal membuat file tujuan: %v", err)
	}
	defer dstFile.Close()

	// Salin konten dari file asli ke file tujuan
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatalf("Gagal menyalin file: %v", err)
	}
}

func SetStatic() {
	cmd := exec.Command("xcopy", "public", "build\\public", "/E", "/I", "/Y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to copy public directory: %v", err)
	}
}

func SetView() {
	cmd := exec.Command("xcopy", "resources\\views", "build\\views", "/E", "/I", "/Y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to copy views directory: %v", err)
	}
}

func SetSQlite() {
	var DBConfig = config.NewDatabaseConfig().Load()

	if DBConfig.DB_CONNECTION == "sqlite" {
		// Membuat folder build\database
		if err := os.Mkdir("build\\database", 0755); err != nil {
			log.Fatalf("Failed to create build directory: %v", err)
		}

		// Menyalin basis data dari database\database.sqlite ke build\database.sqlite
		cmd := exec.Command("xcopy", "database\\database.sqlite", "build\\database", "/Y")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Gagal menyalin basis data: %v", err)
		}
	}
}

func SetLogger() {
	var appConfig = config.NewAppConfig().Load()

	if appConfig.APP_ACTIVE_LOGGING {
		// Membuat folder build\logs
		if err := os.Mkdir("build\\logs", 0755); err != nil {
			log.Fatalf("Failed to create build directory: %v", err)
		}
	}
}

func Build() {
	var appConfig = config.NewAppConfig().Load()

	var filename string = "build\\" + appConfig.APP_NAME + ".exe"
	cmd := exec.Command("go", "build", "-o", filename, "./cmd/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to build application: %v", err)
	}
}
