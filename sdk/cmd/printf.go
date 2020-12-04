package cmd

import (
	"context"
	"fmt"
	"os"
)

// Printf prints to a given context's stdout.
func Printf(ctx context.Context, format string, args ...interface{}) {
	fmt.Fprintf(GetStdout(ctx, os.Stdout), format, args...)
}
