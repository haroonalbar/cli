package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name: "Broskii",
    Usage: "A cli app to support the bros.",
    Action: func(ctx *cli.Context) error {
      fmt.Println("Boii")
      return nil
    },
	}
  if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
