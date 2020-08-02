// package zelda
package main

import (
	"fmt"
	"log"
	"os"
	"zelda/apps/www"

	"github.com/urfave/cli"
)

var banner = `
=======================================================================

	Zelda

>>  Explore Go/Gin And Development

=======================================================================
`

func main() {
	fmt.Println(banner)
	app := cli.NewApp()
	app.Name = "zelda"
	app.Usage = "Explore Go/Gin And Development"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serve",
			Action: func(c *cli.Context) error {
				www.HttpServer()
				return nil
			},
		},
		{
			Name:    "validator",
			Aliases: []string{"v"},
			Usage:   "serve",
			Action: func(c *cli.Context) error {
				www.ValidateForm()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
