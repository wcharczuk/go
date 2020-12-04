package cmd

import "context"

// Warning prints warning message if the verbose flag is enabled.
func Warning(ctx context.Context, args ...interface{}) {
	if !IsVerbose(ctx) {
		return
	}
	Println(ctx, append([]interface{}{"[WARN]"}, args...)...)
}
