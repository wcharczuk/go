package profanity

import "go.charczuk.com/sdk/ex"

// Errors
const (
	ErrFailure          ex.Class = "profanity failure"
	ErrRequired         ex.Class = "a required field is unset"
	ErrContentsRequired ex.Class = "contents rule spec must provide `contains`, `glob` or `regex` values"
)
