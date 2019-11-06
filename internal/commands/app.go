package commands

import (
	"github.com/la3mmchen/marius/internal/types"
	"github.com/urfave/cli"
)

var (
	// RulesPath <tbd>
	RulesPath string
	// OutPath <tbd>
	OutPath string
	// TemplatePath <tbd>
	TemplatePath string
)

// GetApp <tbd>
func GetApp(cfg types.Configuration) *cli.App {
	app := cli.NewApp()
	app.Name = cfg.AppName
	app.Usage = cfg.AppUsage
	app.Version = cfg.AppVersion
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Destination: &RulesPath,
			Value:       cfg.DataPath,
			Usage:       "Path to look up rules.",
		},
		cli.StringFlag{
			Name:        "config, c",
			Destination: &TemplatePath,
			Value:       cfg.TemplatePath,
			Usage:       "Path to retrieve the base template from.",
		},
		cli.StringFlag{
			Name:        "out, o",
			Destination: &OutPath,
			Value:       cfg.OutputPath,
			Usage:       "Path to write the created templates.",
		},
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print app version",
	}

	app.Commands = []cli.Command{
		list(cfg), template(cfg),
	}

	return app
}
