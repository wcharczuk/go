package cmd

import "context"

// Debug prints debugging output if the flag is enabled.
func Debug(ctx context.Context, args ...interface{}) {
	if !IsDebug(ctx) {
		return
	}
	Println(ctx, append([]interface{}{"[DEBUG]"}, args...)...)
}
