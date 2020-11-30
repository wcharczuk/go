package profanity

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"

	"go.charczuk.com/sdk/validation"
)

var (
	_ Rule = (*Contents)(nil)
)

// Errors
var (
	ErrContentsRequired = errors.New("contents rule spec must provide `contains`, `glob` or `regex` values")
)

// Contents creates a new contents rule.
// It failes if any of the expressions match.
type Contents struct {
	// Contains is a filter set that uses `strings.Contains` as the predicate.
	Contains ContainsFilter `yaml:"contains,omitempty"`
	// Glob is a filter set that uses `Glob` as the predicate.
	Glob GlobFilter `yaml:"glob,omitempty"`
	// Regex is a filter set that uses `regexp.MustMatch` as the predicate
	Regex RegexFilter `yaml:"regex,omitempty"`
}

// Validate returns validators.
func (cm Contents) Validate() error {
	if len(cm.Contains.Filter.Include) == 0 && len(cm.Glob.Filter.Include) == 0 && len(cm.Regex.Filter.Include) == 0 {
		return &validation.Error{
			Err: ErrContentsRequired,
		}
	}
	return nil
}

// Check implements Rule.
func (cm Contents) Check(filename string, contents []byte) (result RuleResult) {
	scanner := bufio.NewScanner(bytes.NewReader(contents))

	var line int
	var lineText string
	var containsInclude, containsExclude string
	var globInclude, globExclude string
	var regexInclude, regexExclude string
	for scanner.Scan() {
		line++
		lineText = scanner.Text()

		containsInclude, containsExclude = cm.Contains.Match(lineText)
		globInclude, globExclude = cm.Glob.Match(lineText)
		regexInclude, regexExclude = cm.Regex.Match(lineText)

		if cm.Contains.AllowMatch(containsInclude, containsExclude) ||
			cm.Glob.AllowMatch(globInclude, globExclude) ||
			cm.Regex.AllowMatch(regexInclude, regexExclude) {

			var tokens []string
			if containsInclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents contains include: %q", containsInclude))
			}
			if containsExclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents contains exclude: %q", containsExclude))
			}
			if globInclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents glob include: %q", globInclude))
			}
			if globExclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents glob exclude: %q", globExclude))
			}
			if regexInclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents regex include: %q", regexInclude))
			}
			if regexExclude != "" {
				tokens = append(tokens, fmt.Sprintf("contents regex exclude: %q", regexExclude))
			}
			result = RuleResult{
				File:    filename,
				Line:    line,
				Message: strings.Join(tokens, ", "),
			}
		}
	}

	return RuleResult{OK: true}
}

// String implements fmt.Stringer.
func (cm Contents) String() string {
	var tokens []string
	if len(cm.Contains.Filter.Include) > 0 {
		tokens = append(tokens, fmt.Sprintf("contain: %s", cm.Contains.String()))
	}
	if len(cm.Glob.Filter.Include) > 0 {
		tokens = append(tokens, fmt.Sprintf("glob: %s", cm.Glob.String()))
	}
	if len(cm.Regex.Filter.Include) > 0 {
		tokens = append(tokens, fmt.Sprintf("regex: %s", cm.Glob.String()))
	}
	return fmt.Sprintf("[contents %s]", strings.Join(tokens, ","))
}
