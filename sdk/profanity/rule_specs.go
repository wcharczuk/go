package profanity

import "strings"

// RuleSpecFile is a map of string rule id to rule item.
//
// It is the format for profanity rule files.
type RuleSpecFile map[string]RuleSpec

// Rules returns the
func (rsf RuleSpecFile) Rules() []RuleSpec {
	var rules []RuleSpec
	for id, rule := range rsf {
		rule.ID = id
		rules = append(rules, rule)
	}
	return rules
}

// String implements fmt.Stringer.
func (rsf RuleSpecFile) String() string {
	if len(rsf) == 0 {
		return "<empty>"
	}
	var output []string
	for _, rule := range rsf.Rules() {
		output = append(output, rule.String())
	}
	return strings.Join(output, "\n")
}
