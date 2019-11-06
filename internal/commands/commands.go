package commands

import (
	"fmt"

	"github.com/la3mmchen/marius/internal/parse"
	"github.com/la3mmchen/marius/internal/types"
	"github.com/la3mmchen/marius/internal/write"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func list(cfg types.Configuration) cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "List rules in configured path",
	}

	cmd.Action = func(c *cli.Context) error {
		ruleFiles, err := parse.PopulateFromPath(cfg, RulesPath)

		if err != nil {
			return errors.Wrap(err, "failed to parse config")
		}

		for _, file := range ruleFiles {
			fmt.Printf("%v \n", file.FileName)

			for _, group := range file.Groups {
				fmt.Printf("|- %v \n", group.Name)

				for _, rule := range group.Rules {
					fmt.Printf("|--- %v \n", rule.Alert)
				}
			}

			fmt.Printf("\n \n")
		}
		return nil
	}
	return cmd
}

func template(cfg types.Configuration) cli.Command {
	cmd := cli.Command{
		Name:  "template",
		Usage: "Create test template for in configured path",
	}

	cmd.Action = func(c *cli.Context) error {
		fmt.Printf("Creating templates from rules in [%v] and write to [%v] \n", RulesPath, OutPath)
		ruleFiles, err := parse.PopulateFromPath(cfg, RulesPath)

		for _, file := range ruleFiles {
			outFile := "data/test-" + file.FileName // TODO: switch to real output path
			fmt.Printf("|-- creating %v (source %v) \n", outFile, file.FileName)
			err = write.Template(file, OutPath, TemplatePath, outFile)
		}

		if err != nil {
			return errors.Wrap(err, "failed to parse config")
		}

		return errors.Wrap(err, "failed to parse config")
	}
	return cmd
}
