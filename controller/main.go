package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/clyhs/goweb/controller/commands"
	"github.com/codegangsta/cli"
)

const (
	STORE_KEY = "helloweb"
)

func main() {

	app := cli.NewApp()
	app.Name = "helloweb"
	app.Usage = ""
	app.Author = ""
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "run helloweb controller",
			Action: commands.CmdServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "listen, l",
					Usage: "listen address",
					Value: ":8070",
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
