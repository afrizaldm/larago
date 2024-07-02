package commands

import (
	"log"
	"os"
	"os/exec"
	"simple-api/config"
	"strings"

	"github.com/urfave/cli/v2"
)

var ServeCommand cli.Command = cli.Command{
	Name:  "serve",
	Usage: "Serve the Go application",
	Action: func(c *cli.Context) error {

		appConfig := config.NewAppConfig().Load()
		port := strings.ReplaceAll(appConfig.APP_PORT, ":", "")
		os.Setenv("PORT", port)
		os.Setenv("GIN_MODE", "debug")

		// Menjalankan aplikasi dalam mode debug
		cmd := exec.Command("go", "run", "./cmd/main.go")
		// cmd := exec.Command("gin", "--path", ".", "--build", "cmd", "--port", "3000")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to serve application: %v", err)
		}

		return nil
	},
}
