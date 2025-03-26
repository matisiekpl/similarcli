package main

import (
	"fmt"
	"os"

	"github.com/matisiekpl/similarcli/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	app := &cli.App{
		Name:      "similarcli",
		Usage:     "CLI tool for fetching website analytics from SimilarWeb",
		UsageText: "similarcli [command options] DOMAIN",
		ArgsUsage: "DOMAIN",
		Authors: []*cli.Author{
			{
				Name: "Mateusz Woźniak",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				cli.ShowAppHelp(c)
				return nil
			}
			if c.NArg() != 1 {
				return fmt.Errorf("only domain argument is required")
			}
			similarWebService := service.NewSimilarWebService()
			return similarWebService.Print(c.Args().First())
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
