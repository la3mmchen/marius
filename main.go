package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/la3mmchen/marius/internal/commands"
	"github.com/la3mmchen/marius/internal/types"
)

var (
	configFile string = "config.json"
	// AppVersion Version of the app. Must be injected during the build.
	AppVersion string
)

func main() {
	var Cfg = types.Configuration{
		AppUsage:     "Create templates for unit test from existing prometheus alert rules.",
		AppName:      "marius",
		AppVersion:   AppVersion,
		DataPath:     "data",
		TemplatePath: "config/template.tpl",
		OutputPath:   "data/tests",
		Debug:        false,
	}

	// load config if it is present
	if _, err := os.Stat(filepath.Join(".", configFile)); !os.IsNotExist(err) {
		file, err := os.Open(filepath.Join(".", configFile))
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Cfg)
		if err != nil {
			os.Exit(1)
		}
	}

	app := commands.GetApp(Cfg)

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
