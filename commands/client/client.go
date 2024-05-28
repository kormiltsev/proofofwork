package client

import (
	"net/url"

	"github.com/urfave/cli"

	"github.com/kormiltsev/proofofwork/client/app"
	"github.com/kormiltsev/proofofwork/config"
)

// Runs a client part of application
var Command = cli.Command{
	Name:  "client",
	Usage: "Request the Words of Wisdom",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "url",
			Usage:       "Target URL.",
			EnvVar:      "POW_CLIENT_URL",
			Value:       "http://localhost:12000/words",
			Destination: &config.URL,
		},
		cli.BoolFlag{
			Name:   "endless",
			Usage:  "Run requests endless.",
			EnvVar: "POW_CLIENT_ENDLESS",
		},
	},
	Action: func(c *cli.Context) error {
		curl := c.String("url")
		_, err := url.ParseRequestURI(curl)
		if err != nil {
			return err
		}

		if c.Bool("endless") {
			return app.Endless(curl)
		}

		return app.OneRequest(curl)
	},
}
