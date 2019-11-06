package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/la3mmchen/marius/internal/commands"
	"github.com/la3mmchen/marius/internal/types"
)

var (
	configFile string = "config/config.json"
	GitCommit  string
)

func main() {
	var Cfg types.Configuration

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// load config
	file, err := os.Open(configFile)
	if err != nil {
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil {
		os.Exit(1)
	}

	Cfg.AppName = filepath.Base((dir))
	Cfg.AppVersion = GitCommit

	app := commands.GetApp(Cfg)

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
