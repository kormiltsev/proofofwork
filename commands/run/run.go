package run

import (
	"context"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/tomb.v2"

	"github.com/kormiltsev/proofofwork/api"
	"github.com/kormiltsev/proofofwork/config"
)

// Runs server
var Command = cli.Command{
	Name:  "run",
	Usage: "Run the Words of Wisdom API",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "socket",
			Usage:       "REST API `socket` string",
			EnvVar:      "POW_SOCKET",
			Value:       ":8080",
			Destination: &config.Host,
		},
		cli.IntFlag{
			Name:        "initial-difficulty",
			Usage:       "Server will starts with this difficulty",
			EnvVar:      "POW_INITIAL_DIFFICULTY",
			Value:       16,
			Destination: &config.Difficulty,
		},
		cli.IntFlag{
			Name:        "cache-limit",
			Usage:       "Limits size of requests kept in memory",
			EnvVar:      "POW_CACHE_LIMIT",
			Value:       1000,
			Destination: &config.Limit,
		},
	},
	Action: func(c *cli.Context) error {
		t, ctx := tomb.WithContext(context.Background())

		httpListener, err := listenerFromSocket(c.String("socket"))
		if err != nil {
			return err
		}

		// start API
		return api.Run(ctx, t, httpListener)
	},
}

func listenerFromSocket(socket string) (lis net.Listener, err error) {
	if strings.HasPrefix(socket, "unix://") {
		f := strings.TrimPrefix(socket, "unix://")
		if _, checkErr := os.Stat(f); checkErr == nil {
			err = os.Remove(f)
			if err != nil {
				return
			}
		}
		if lis, err = net.Listen("unix", f); err == nil {
			err = os.Chmod(f, 0o600)
			if err != nil {
				return
			}
		}

	} else {
		lis, err = net.Listen("tcp", strings.TrimPrefix(socket, "tcp://"))
		if err != nil {
			return
		}
	}
	return
}
