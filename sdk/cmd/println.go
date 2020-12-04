package cmd

import (
	"context"
	"fmt"
	"os"
)

// Println prints to a given context's stdout.
func Println(ctx context.Context, args ...interface{}) {
	fmt.Fprintln(GetStdout(ctx, os.Stdout), args...)
}
