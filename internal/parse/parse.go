package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/la3mmchen/marius/internal/types"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// PopulateFromPath <tbd>
func PopulateFromPath(config types.Configuration, path string) ([]types.RuleFile, error) {
	var ruleFiles []types.RuleFile
	skipPatterns := []string{"^\\.", "test-.*", ".*\\.tpl"}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		var file types.RuleFile
		var usable = true

		// skip if file points to dir
		if f.IsDir() {
			usable = false
			continue
		}

		// handle custom skipPatterns
		for _, pattern := range skipPatterns {
			matched, _ := regexp.MatchString(pattern, f.Name())
			if matched {
				usable = false
				break
			}
		}

		// only work on files deemed as usable
		if config.Debug {
			fmt.Printf("$ file: [%v] (usable: [%v]) \n", f.Name(), usable)
		}

		if usable {
			// Initial parsing from the files
			file, err = loadConfig(path, f.Name())
			file.FileName = f.Name()

			file.Metrics = extractCustomMetrics(file)
			file.Labels, _ = extractCustomLabels(file)
			file.TmplSeries = permuatateMetricLabels(file) // build any useful combination for template section `input_series`

			ruleFiles = append(ruleFiles, file)
		}

		if err != nil {
			return ruleFiles, errors.Wrap(err, "failed to load configs.")
		}
	}
	return ruleFiles, err
}

func loadConfig(path string, file string) (types.RuleFile, error) {
	var config types.RuleFile

	f, err := os.Open(path + "/" + file)
	defer f.Close()
	if err != nil {
		return types.RuleFile{}, errors.Wrap(err, "failed to open file")
	}

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return types.RuleFile{}, errors.Wrap(err, "failed to parse config")
	}

	return config, err
}

func extractCustomMetrics(in types.RuleFile) []string {
	var out []string
	artefacts := []string{"{", "=="} // TODO: make the regex not contain some control characters liste here

	for _, group := range in.Groups {
		for _, rule := range group.Rules {

			if !(strings.Contains(rule.Expr, "{")) {
				var pattern = regexp.MustCompile(`(?P<Metric>[a-zA-Z_:][a-zA-Z0-9_:][^{}]*)\w?==`)
				for _, item := range pattern.FindAllString(rule.Expr, -1) {

					// remove control chars
					for _, x := range artefacts {
						if strings.Contains(item, x) {
							item = removeChar(item, x)
						}
					}
					item = strings.TrimSpace(item)
					out = append(out, item)
				}
			} else {
				// extract metrics that defines labels
				var pattern = regexp.MustCompile(`(?m)(?P<Metric>[a-zA-Z_:][a-zA-Z0-9_:]*)\{`)

				for _, item := range pattern.FindAllString(rule.Expr, -1) {

					// remove control chars
					for _, x := range artefacts {
						if strings.Contains(item, x) {
							item = removeChar(item, x)
						}
					}

					out = append(out, item)
				}
			}

		}
	}

	return out
}

func removeChar(in string, remove string) string {
	return strings.Replace(in, remove, "", -1)
}

// Extract Labels that are used within Annotation and Summaries
func extractCustomLabels(in types.RuleFile) (map[string]int, types.Annotation) {
	var labels map[string]int
	labels = make(map[string]int)

	var annotations types.Annotation

	for _, group := range in.Groups {
		for _, rule := range group.Rules {

			re := regexp.MustCompile(`{{\s+\$labels\.(?P<LabelName>[a-zA-Z_:][a-zA-Z0-9_:]*)\s+}}`)
			groupNames := re.SubexpNames()
			for _, match := range re.FindAllStringSubmatch(rule.Annotations.Summary, -1) {
				for groupIdx, matchgroup := range match {
					name := groupNames[groupIdx]
					if name == "" { // skip the non-named match
						continue
					}
					matchgroup = strings.TrimSpace(matchgroup)
					labels[matchgroup] = 1
				}
			}
		}
	}
	return labels, annotations
}

func permuatateMetricLabels(in types.RuleFile) map[string]string {
	out := make(map[string]string)
	for _, metric := range in.Metrics {
		for label := range in.Labels {
			out[metric] = label
		}
	}
	return out
}
