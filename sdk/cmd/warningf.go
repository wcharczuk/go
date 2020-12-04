package cmd

import "context"

// Warningf prints warning message if the verbose flag is enabled.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	if !IsVerbose(ctx) {
		return
	}
	Printf(ctx, "[WARN] "+format+"\n", args...)
}
