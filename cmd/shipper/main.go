package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/shipper/pkg/config"
	"github.com/sudhanshuraheja/shipper/pkg/logger"
)

func main() {
	config.Init()
	logger.Init()

	logger.Infoln("Shipper")
	Init()
}

// Init : start the cli wrapper
func Init() *cli.App {
	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "This service ships binaries to the server"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	return app
}
