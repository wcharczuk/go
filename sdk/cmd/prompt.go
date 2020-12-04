package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

// Prompt gives a prompt and reads input until newlines from a given set of streams.
func Prompt(ctx context.Context, prompt string) string {
	fmt.Fprint(GetStdout(ctx, os.Stdout), prompt)

	scanner := bufio.NewScanner(GetStdin(ctx, os.Stdin))
	var output string
	if scanner.Scan() {
		output = scanner.Text()
	}
	return output
}
