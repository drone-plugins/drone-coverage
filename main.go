package main

import (
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "coverage generator"
	app.Usage = "generate and publish coverage reports"
	app.Version = version
	app.Commands = []cli.Command{
		LcovCmd,
		GocovCmd,
		PublishCmd,
	}

	app.Run(os.Args)
}
