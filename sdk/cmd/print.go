package cmd

import (
	"context"
	"fmt"
	"os"
)

// Print prints to a given context's stdout.
func Print(ctx context.Context, args ...interface{}) {
	fmt.Fprint(GetStdout(ctx, os.Stdout), args...)
}
