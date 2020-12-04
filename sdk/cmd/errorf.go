package cmd

import (
	"context"
	"fmt"
	"os"
)

// Errorf prints to a given context's stderr.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	fmt.Fprintf(GetStderr(ctx, os.Stderr), format, args...)
}
