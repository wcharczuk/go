package cmd

import (
	"context"
	"fmt"
	"os"
)

// Error prints to a given context's stderr.
func Error(ctx context.Context, args ...interface{}) {
	fmt.Fprint(GetStderr(ctx, os.Stderr), args...)
}
