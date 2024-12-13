package cmd

import (
	cli "github.com/urfave/cli/v2"

	"github.com/incu6us/weather-cli/app"
)

func run() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "run the weather-cli application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "country",
				Value: "Ukraine",
				Usage: "choose country",
			},
			&cli.StringFlag{
				Name:  "city",
				Value: "Kyiv",
				Usage: "choose city",
			},
		},
		Action: func(c *cli.Context) error {
			application, err := app.NewApplication(c.Context, c.String("config"), c.Bool("debug"))
			if err != nil {
				return err
			}
			err = application.Run(c.Context, c.String("country"), c.String("city"))
			if err != nil {
				return err
			}
			return nil
		},
	}
}
