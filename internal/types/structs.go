package types

// Configuration <tbd>
type Configuration struct {
	AppUsage     string
	AppName      string
	AppVersion   string
	DataPath     string
	TemplatePath string
	OutputPath   string
	Debug        bool
}

// RuleFile <tbd>
type RuleFile struct {
	FileName   string
	Metrics    []string
	Labels     map[string]int
	Groups     []Group `yaml:"groups"`
	TmplSeries map[string]string
}

// Group <tbd>
type Group struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

// Rule <tbd>
type Rule struct {
	Alert       string            `yaml:"alert,omitempty"`
	Expr        string            `yaml:"expr,omitempty"`
	For         string            `yaml:"for,omitempty"`
	AllLabels   map[string]string `yaml:"labels,omitempty"`
	Annotations Annotation        `yaml:"annotations"`
}

// Annotation <tbd>
type Annotation struct {
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
}
