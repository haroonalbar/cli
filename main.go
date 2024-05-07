package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name: "Broskii",
    Usage: "A cli app to support the bros.",
	}
	app.Run(os.Args)
}
