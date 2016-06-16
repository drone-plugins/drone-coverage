package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var build string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "coverage"
	app.Usage = "coverage plugin"
	app.Action = run
	app.Version = "1.0.0+" + build
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "repo.fullname",
			Usage:  "repository fullname",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "netrc.username",
			Usage:  "netrc username",
			EnvVar: "DRONE_NETRC_USERNAME",
		},

		// cli.StringFlag{
		// 	Name:   "path",
		// 	Usage:  "git clone path",
		// 	EnvVar: "DRONE_WORKSPACE",
		// },
		// cli.StringFlag{
		// 	Name:   "sha",
		// 	Usage:  "git commit sha",
		// 	EnvVar: "DRONE_COMMIT_SHA",
		// },
		// cli.StringFlag{
		// 	Name:   "ref",
		// 	Value:  "refs/heads/master",
		// 	Usage:  "git commit ref",
		// 	EnvVar: "DRONE_COMMIT_REF",
		// },
		// cli.StringFlag{
		// 	Name:   "event",
		// 	Value:  "push",
		// 	Usage:  "build event",
		// 	EnvVar: "DRONE_BUILD_EVENT",
		// },
		// cli.StringFlag{
		// 	Name:   "netrc.machine",
		// 	Usage:  "netrc machine",
		// 	EnvVar: "DRONE_NETRC_MACHINE",
		// },
		// cli.StringFlag{
		// 	Name:   "netrc.username",
		// 	Usage:  "netrc username",
		// 	EnvVar: "DRONE_NETRC_USERNAME",
		// },
		// cli.StringFlag{
		// 	Name:   "netrc.password",
		// 	Usage:  "netrc password",
		// 	EnvVar: "DRONE_NETRC_PASSWORD",
		// },

		cli.StringFlag{
			Name:   "server",
			Usage:  "server address",
			EnvVar: "PLUGIN_SERVER",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "server token",
			EnvVar: "PLUGIN_TOKEN",
		},
		cli.IntFlag{
			Name:   "threshold",
			Usage:  "",
			EnvVar: "PLUGIN_THRESHOLD",
		},
		cli.StringFlag{
			Name:   "include",
			Usage:  "",
			EnvVar: "PLUGIN_INCLUDE",
		},
		cli.BoolFlag{
			Name:   "must-increase",
			Usage:  "",
			EnvVar: "PLUGIN_MUST_INCREASE",
		},
		cli.StringFlag{
			Name:   "ca-cert",
			Usage:  "",
			EnvVar: "PLUGIN_CA_CERT",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repo: Repo{
			FullName: c.String("repo.fullname"),
		},
		Build: Build{
			Link: c.String("build.link"),
		},
		Netrc: Netrc{
			Login: c.String("netrc.username"),
		},
		Config: Config{
			Server:       c.String("server"),
			Token:        c.String("token"),
			Threshold:    c.Int("threshold"),
			Include:      c.String("include"),
			MustIncrease: c.Bool("must-increase"),
			CACert:       c.String("ca-cert"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
