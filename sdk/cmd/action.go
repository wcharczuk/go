package cmd

import "context"

// Action is a delegate type that represents
// a step in a given target.
type Action func(context.Context) error
