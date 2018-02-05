package main

import (
	"os"
	"sort"

	"github.com/urfave/cli"

	"source.golabs.io/core/shipper/pkg/appcontext"
	"source.golabs.io/core/shipper/pkg/config"
	"source.golabs.io/core/shipper/pkg/logger"
	"source.golabs.io/core/shipper/pkg/uploader"
)

func main() {
	config := config.NewConfig()
	logger := logger.NewLogger(config, os.Stdout)
	ctx := appcontext.NewAppContext(config, logger)

	app := cli.NewApp()
	app.Name = "shipper"
	app.Version = "0.0.1"
	app.Usage = "this binary uploads builds for distribution"

	service := uploader.NewService(ctx)

	app.Action = func(c *cli.Context) error {
		err := service.Upload(c.String("bundle"), c.String("file"))
		if err != nil {
			logger.Infoln(err)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "install the service",
			Action: func(c *cli.Context) error {
				return service.Install(c.String("server"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server, s",
					Usage: "base url of the server",
				},
			},
		},
		{
			Name:  "uninstall",
			Usage: "uninstall the service",
			Action: func(c *cli.Context) error {
				return service.Uninstall()
			},
		},
	}

	app.Flags = []cli.Flag{
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

}
