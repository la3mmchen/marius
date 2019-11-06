package write

import (
	"os"
	"text/template"

	"github.com/la3mmchen/marius/internal/types"
	"github.com/pkg/errors"
)

// Template <tbd>
func Template(ruleFile types.RuleFile, outPath string, tplPath string, outFile string) error {
	// yolo create dir ignore errors \o/
	os.MkdirAll(outPath, os.ModePerm)

	if _, err := os.Stat(tplPath); os.IsNotExist(err) {
		return errors.Wrap(err, "failed to parse config")
	}

	w, err := os.Create(outFile)

	if err != nil {
		return errors.Wrap(err, "unable to create test template file.")
	}

	tpl, err := template.ParseFiles(tplPath)

	_ = tpl.Execute(w, ruleFile)

	if err != nil {
		return errors.Wrap(err, "failed to create template.")
	}

	return nil
}
