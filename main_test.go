package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/la3mmchen/marius/internal/parse"
	"github.com/la3mmchen/marius/internal/types"
)

var (
	dir = "tmpGoTest"
)

func cleanupTestArtefacts(d string) error {
	os.RemoveAll(d)

	return nil
}

func prepareYaml(d string) error {
	os.MkdirAll(d, os.ModePerm)

	err := ioutil.WriteFile(filepath.Join(d, "test.yaml"), []byte("---"), 0755)

	if err != nil {
		return err
	}

	return nil
}

func prepareIgnored(d string) error {
	os.MkdirAll(d, os.ModePerm)

	err := ioutil.WriteFile(filepath.Join(d, "test.tpl"), []byte("random"), 0755)

	if err != nil {
		return err
	}

	return nil
}
func prepareOther(d string) error {
	os.MkdirAll(d, os.ModePerm)

	err := ioutil.WriteFile(filepath.Join(d, "test.file"), []byte("random"), 0755)

	if err != nil {
		return err
	}

	return nil
}

// TestPopulateFromPath We test if we can drop different files (valid, to be ignored and non-valid yaml) into the PopulateFromPath
func TestPopulateFromPath(t *testing.T) {
	cleanupTestArtefacts(dir)

	// Throw in valid yaml
	_ = prepareYaml(dir)

	_, noErr := parse.PopulateFromPath(types.Configuration{}, dir)

	if noErr != nil {
		cleanupTestArtefacts(dir)
		t.Error()
	}

	cleanupTestArtefacts(dir)

	// Throw some file to be ignored
	_ = prepareIgnored(dir)

	_, noErr = parse.PopulateFromPath(types.Configuration{}, dir)

	if noErr != nil {
		cleanupTestArtefacts(dir)
		t.Error()
	}

	// Throw some not to be ignored. PopulateFromPath should return an Error.
	_ = prepareOther(dir)

	_, err := parse.PopulateFromPath(types.Configuration{}, dir)

	if err == nil {
		cleanupTestArtefacts(dir)
		t.Error()
	}

	cleanupTestArtefacts(dir)
}
