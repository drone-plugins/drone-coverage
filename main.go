package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "coverage generator"
	app.Usage = "generate and publish coverage reports"
	app.Version = version
	app.Before = setup
	app.Flags = []cli.Flag{
		cli.StringFlag{
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

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func setup(c *cli.Context) error {
	logrus.SetOutput(os.Stderr)

	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	return nil
}
