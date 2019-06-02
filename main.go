package main

import (
	"fmt"
	"github.com/twocucao/thanos/web"
	"github.com/urfave/cli"
	"log"
	"os"
)

var banner = `
=======================================================================

  _______ _                           
 |__   __| |                          
    | |  | |__   __ _ _ __   ___  ___ 
    | |  | '_ \ / _' | '_ \ / _ \/ __|
    | |  | | | | (_| | | | | (_) \__ \
    |_|  |_| |_|\__,_|_| |_|\___/|___/


>>  Explore Go/Gin And Development

=======================================================================
`

func main() {
	fmt.Println(banner)
	app := cli.NewApp()
	app.Name = "thanos"
	app.Usage = "Explore Go/Gin And Development"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serve",
			Action: func(c *cli.Context) error {
				web.HttpServer()
				return nil
			},
		},
		{
			Name:    "validator",
			Aliases: []string{"v"},
			Usage:   "serve",
			Action: func(c *cli.Context) error {
				web.ValidateForm()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
