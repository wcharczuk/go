package cmd

import (
	"context"
	"fmt"
	"os"
)

// Errorln prints to a given context's stderr.
func Errorln(ctx context.Context, args ...interface{}) {
	fmt.Fprintln(GetStderr(ctx, os.Stderr), args...)
}
