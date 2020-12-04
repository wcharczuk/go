package profanity

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go.charczuk.com/sdk/ex"
	"gopkg.in/yaml.v3"
)

// New creates a new profanity engine with a given set of config options.
func New(options ...ConfigOption) *Profanity {
	var cfg Config
	for _, option := range options {
		option(&cfg)
	}
	return &Profanity{
		Config: cfg,
	}
}

// Profanity parses rules from the filesystem and applies them to a given root path.
// Creating a full rules set.
type Profanity struct {
	Config Config
	Stdout io.Writer
	Stderr io.Writer
}

// Process processes the profanity rules.
func (p *Profanity) Process() error {
	p.Verbosef("using rules file: %q", p.Config.RulesFileOrDefault())
	if fileFilter := p.Config.Files.String(); fileFilter != "" {
		p.Verbosef("using file filter: %s", fileFilter)
	}
	if dirFilter := p.Config.Dirs.String(); dirFilter != "" {
		p.Verbosef("using dir filter: %s", dirFilter)
	}
	err := p.Walk(p.Config.PathOrDefault())
	if err != nil {
		if err != ErrFailure {
			return err
		}
		p.Verbosef("profanity %s!", ansi.Red("failed"))
		return nil
	}
	p.Verbosef("profanity %s!", ansi.Green("ok"))
	return nil
}

// Walk walks a given path, inheriting a set of rules.
func (p *Profanity) Walk(path string, rules ...RuleSpec) error {
	dirs, files, err := ListDir(path)
	if err != nil {
		return ex.New("profanity; invalid walk path", ex.OptMessagef("path: %q", path), ex.OptInner(err))
	}

	var didFail bool
	var fullFilePath string
	for _, file := range files {
		if file.Name() == p.Config.RulesFileOrDefault() {
			fullFilePath = filepath.Join(path, file.Name())
			p.Debugf("reading rules file: %q", filepath.Join(path, fullFilePath))
			foundRules, err := p.ReadRuleSpecsFile(fullFilePath)
			if err != nil {
				return err
			}
			rules = append(rules, foundRules...)
		}
	}

	for _, file := range files {
		if file.Name() == p.Config.RulesFileOrDefault() {
			continue
		}

		fullFilePath = filepath.Join(path, file.Name())
		if p.Config.Files.Allow(fullFilePath) {
			contents, err := ioutil.ReadFile(fullFilePath)
			if err != nil {
				return err
			}
			for _, rule := range rules {
				if rule.Files.Allow(fullFilePath) {
					res := rule.Check(fullFilePath, contents)
					if res.Err != nil {
						return res.Err
					}
					if !res.OK {
						didFail = true
						p.Errorf("%v\n", p.FormatRuleResultFailure(rule, res))
						if p.Config.FailFastOrDefault() {
							return ErrFailure
						}
					}
				}
			}
		}
	}

	var fullDirPath string
	for _, dir := range dirs {
		if dir.Name() == ".git" {
			continue
		}
		if strings.HasPrefix(dir.Name(), "_") {
			continue
		}
		fullDirPath = filepath.Join(path, dir.Name())
		if p.Config.Dirs.Allow(fullDirPath) {
			if err := p.Walk(fullDirPath, rules...); err != nil {
				if err != ErrFailure || p.Config.FailFastOrDefault() {
					return err
				}
				didFail = true
			}
		}
	}

	if didFail {
		return ErrFailure
	}
	return nil
}

// ReadRuleSpecsFile reads rules from a file path.
//
// It is expected to be passed the fully qualified path for the rules file.
func (p *Profanity) ReadRuleSpecsFile(filename string) (rules []RuleSpec, err error) {
	contents, readErr := os.Open(filename)
	if readErr != nil {
		err = ex.New(readErr, ex.OptMessagef("file: %s", filename))
		return
	}
	defer contents.Close()
	rules, err = p.ReadRuleSpecsFromReader(filename, contents)
	return
}

// ReadRuleSpecsFromReader reads rules from a reader.
func (p *Profanity) ReadRuleSpecsFromReader(filename string, reader io.Reader) (rules []RuleSpec, err error) {
	fileRules := make(RuleSpecFile)
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)
	yamlErr := decoder.Decode(&fileRules)
	if yamlErr != nil {
		err = ex.New("cannot unmarshal rules file", ex.OptMessagef("file: %s", filename), ex.OptInnerClass(yamlErr))
		return
	}
	for _, rule := range fileRules.Rules() {
		rule.SourceFile = filename
		if validationErr := rule.Validate(); validationErr != nil {
			p.Debugf("rule file %q fails validation", filename)
			err = validationErr
			return
		}
		rules = append(rules, rule)
	}
	return
}

// FormatRuleResultFailure formats a rule result with the rule that produced it.
func (p Profanity) FormatRuleResultFailure(r RuleSpec, rr RuleResult) error {
	if rr.OK {
		return nil
	}
	var lines []string
	lines = append(lines, fmt.Sprintf("%s:%d", ansi.Bold(ansi.ColorWhite, rr.File), rr.Line))
	lines = append(lines, fmt.Sprintf("\t%s: %s", ansi.LightBlack("id"), r.ID))
	if r.Description != "" {
		lines = append(lines, fmt.Sprintf("\t%s: %s", ansi.LightBlack("description"), r.Description))
	}
	lines = append(lines, fmt.Sprintf("\t%s: %s", ansi.LightBlack("status"), ansi.Red("failed")))
	lines = append(lines, fmt.Sprintf("\t%s: %s", ansi.LightBlack("rule"), rr.Message))
	return fmt.Errorf(strings.Join(lines, "\n"))
}

// Verbosef prints a verbose message.
func (p *Profanity) Verbosef(format string, args ...interface{}) {
	if p.Config.VerboseOrDefault() {
		p.Printf("[VERBOSE] "+format+"\n", args...)
	}
}

// Debugf prints a debug message.
func (p *Profanity) Debugf(format string, args ...interface{}) {
	if p.Config.DebugOrDefault() {
		p.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// Printf writes to the output stream.
func (p *Profanity) Printf(format string, args ...interface{}) {
	if p.Stdout != nil {
		fmt.Fprintf(p.Stdout, format, args...)
	}
}

// Errorf writes to the error output stream.
func (p *Profanity) Errorf(format string, args ...interface{}) {
	if p.Stderr != nil {
		fmt.Fprintf(p.Stderr, format, args...)
	}
}
