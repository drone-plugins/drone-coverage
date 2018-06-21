package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "drone-coverage plugin"
	app.Usage = "generate and publish coverage reports"
	app.Version = version
	app.Before = setup
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "ci",
			Usage:  "continuous integration mode",
			EnvVar: "CI",
		},
		cli.BoolTFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},
	}
	app.Commands = []cli.Command{
		LcovCmd,
		GocovCmd,
		PublishCmd,
	}

	log.WithFields(log.Fields{
		"version": app.Version,
	}).Info(app.Name)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func setup(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	return nil
}
