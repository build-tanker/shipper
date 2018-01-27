package main

import (
	"os"
	"sort"

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
	app.Action = func(c *cli.Context) error {
		logger.Infoln("Getting ready to ship")
		logger.Infoln(c.String("key"))
		logger.Infoln(c.String("file"))
		logger.Infoln(c.String("bundle"))
		return nil
	}

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Usage: "access key for authentication",
		},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "file to be uploaded",
		},
		cli.StringFlag{
			Name:  "bundle, b",
			Usage: "app bundle to link to",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	return app
}
