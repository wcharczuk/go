package cmd

import "context"

// Debugf prints debugging output if the flag is enabled.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	if !IsDebug(ctx) {
		return
	}
	Printf(ctx, "[DEBUG] "+format+"\n", args...)
}
