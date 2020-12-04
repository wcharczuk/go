package cmd

import "context"

// Verbose prints debugging output if the flag is enabled.
func Verbose(ctx context.Context, args ...interface{}) {
	if !IsVerbose(ctx) {
		return
	}
	Println(ctx, args...)
}
