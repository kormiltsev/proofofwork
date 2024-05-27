package main

import (
	"os"

	"github.com/inconshreveable/log15"
	"github.com/urfave/cli"

	"github.com/kormiltsev/proofofwork/commands/client"
	"github.com/kormiltsev/proofofwork/commands/run"
	"github.com/kormiltsev/proofofwork/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "words-of-wisdom"
	app.Usage = "REST API server and a client."
	app.Version = version.Version + " (" + version.GitCommit + ")"
	app.Before = func(c *cli.Context) error {
		log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlDebug, log15.CallerFileHandler(log15.StreamHandler(os.Stdout, log15.LogfmtFormat()))))
		return nil
	}
	app.Commands = []cli.Command{
		run.Command,
		client.Command,
	}

	if err := app.Run(os.Args); err != nil {
		log15.Error("done", log15.Ctx{
			"err": err,
		})
	}
}
