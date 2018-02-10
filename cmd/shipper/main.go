package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/shipper/pkg/appcontext"
	"github.com/sudhanshuraheja/shipper/pkg/config"
	"github.com/sudhanshuraheja/shipper/pkg/logger"
	"github.com/sudhanshuraheja/shipper/pkg/uploader"
)

func main() {
	config := config.NewConfig([]string{"$HOME"})
	logger := logger.NewLogger(config, os.Stdout)
	ctx := appcontext.NewAppContext(config, logger)

	app := cli.NewApp()
	app.Name = "shipper"
	app.Version = "0.0.1"
	app.Usage = "this binary uploads builds for distribution"

	service := uploader.NewService(ctx)

	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "install the service, as, shipper install --server http://public.betas.in",
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
			Usage: "uninstall the service, as, shipper uninstall",
			Action: func(c *cli.Context) error {
				return service.Uninstall()
			},
		},
		{
			Name:  "upload",
			Usage: "upload a file to the service, as, shipper upload --bundle com.me.app --file ./file.ipa",
			Action: func(c *cli.Context) error {
				return service.Upload(c.String("bundle"), c.String("file"))
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "file to be uploaded",
				},
				cli.StringFlag{
					Name:  "bundle, b",
					Usage: "app bundle to link to",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
