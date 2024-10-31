package app

import (
	"go-cli/app/commands"

	"github.com/urfave/cli"
)

// Return CLI ready for execution
func Generate() *cli.App {
	commands := []cli.Command{
		commands.IP(),
		commands.ServerName(),
	}

	return &cli.App{
		Name:     "GO CLI: IP Seeker",
		Usage:    "Seek for IP and server names.",
		Commands: commands,
	}
}
