package commands

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

var ConsoleClearCommand cli.Command = cli.Command{
	Name:  "console:clear",
	Usage: "Clear the console",
	Action: func(c *cli.Context) error {
		var cmd *exec.Cmd

		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		default:
			cmd = exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		// cmd.Stdout = nil // Ganti dengan output stdout yang sesuai jika ada
		// cmd.Stderr = nil // Ganti dengan output stderr yang sesuai jika ada
		// if err := cmd.Run(); err != nil {
		// 	log.Fatalf("Failed to clear console: %v", err)
		// }

		return nil
	},
}
