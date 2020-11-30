package profanity

import "errors"

// Defaults
const (
	DefaultPath      = "."
	DefaultRulesFile = ".profanity.yml"
)

// Glob constants
const (
	Star = "*"
	Root = "."

	GoFiles     = "*.go"
	GoTestFiles = "*_test.go"
)

// Errors
var (
	ErrFailure = errors.New("profanity failure")
)
