package cmd

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

// Exec forks a given shell command and arguments, suppressing output.
//
// The environment from the context is passed to the command, if you
// want to remove the environment, set the context environment to an
// empty variable set.
func Exec(ctx context.Context, command string, args ...string) error {
	if len(args) > 0 {
		Debugf(ctx, "exec: %s %s", command, strings.Join(args, " "))
	} else {
		Debugf(ctx, "exec: %s", command)
	}
	commandResolved, err := exec.LookPath(command)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, commandResolved, args...)
	cmd.Env = GetEnviron(ctx, os.Environ())
	return cmd.Run()
}
