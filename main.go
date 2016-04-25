package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "coverage generator"
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

	app.Run(os.Args)
}

func setup(c *cli.Context) error {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	logrus.SetOutput(os.Stderr)

	return nil
}
