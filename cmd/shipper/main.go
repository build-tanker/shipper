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
	logger := logger.NewLogger(config)
	ctx := appcontext.NewAppContext(config, logger)

	logger.Infoln("Starting shipper")

	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this binary uploads builds for distribution"

	uploader := uploader.NewUploader(ctx)

	app.Action = func(c *cli.Context) error {
		err := uploader.Upload(c.String("bundle"), c.String("file"))
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
				return uploader.Install()
			},
		},
		{
			Name:  "uninstall",
			Usage: "uninstall the service",
			Action: func(c *cli.Context) error {
				return uploader.Uninstall()
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
