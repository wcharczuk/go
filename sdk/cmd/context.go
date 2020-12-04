package cmd

import (
	"context"
	"io"
)

type argsKey struct{}

// WithArgs sets the envionment variables for a given context.
func WithArgs(ctx context.Context, args []string) context.Context {
	return context.WithValue(ctx, argsKey{}, args)
}

// GetArgs gets the arguments for a given context.
// Note that the arguments will be everything *not* captured by the command.
// Example: repoctl foo bar, where "foo" is the command, args will be []string{"bar"}.
func GetArgs(ctx context.Context) []string {
	if value := ctx.Value(argsKey{}); value != nil {
		if typed, ok := value.([]string); ok {
			return typed
		}
	}
	return nil
}

type targetsKey struct{}

// WithTargets sets the command targets for a given context.
func WithTargets(ctx context.Context, targets Targets) context.Context {
	return context.WithValue(ctx, targetsKey{}, targets)
}

// GetTargets gets the command targets available.
func GetTargets(ctx context.Context) Targets {
	if value := ctx.Value(targetsKey{}); value != nil {
		if typed, ok := value.(Targets); ok {
			return typed
		}
	}
	return nil
}

// GetTarget gets a target by name from a given context.
func GetTarget(ctx context.Context, name string) (Target, bool) {
	targets := GetTargets(ctx)
	if targets == nil {
		return nil, false
	}
	target, ok := targets[name]
	return target, ok
}

type environKey struct{}

// WithEnviron sets the envionment variables for a given context.
func WithEnviron(ctx context.Context, environ []string) context.Context {
	return context.WithValue(ctx, environKey{}, environ)
}

// GetEnviron gets the environment variables for a given context or default.
func GetEnviron(ctx context.Context, defaultEnviron []string) []string {
	if value := ctx.Value(environKey{}); value != nil {
		if typed, ok := value.([]string); ok {
			return typed
		}
	}
	return defaultEnviron
}

type stdoutKey struct{}

// WithStdout sets the stdout writer on a given context.
func WithStdout(ctx context.Context, writer io.Writer) context.Context {
	return context.WithValue(ctx, stdoutKey{}, writer)
}

// GetStdout gets a stdout writer from a context, or a default.
func GetStdout(ctx context.Context, defaultWriter io.Writer) io.Writer {
	if value := ctx.Value(stdoutKey{}); value != nil {
		if typed, ok := value.(io.Writer); ok {
			return typed
		}
	}
	return defaultWriter
}

type stderrKey struct{}

// WithStderr sets the stderr writer on a given context.
func WithStderr(ctx context.Context, writer io.Writer) context.Context {
	return context.WithValue(ctx, stdoutKey{}, writer)
}

// GetStderr gets a stderr writer from a context, or a default.
func GetStderr(ctx context.Context, defaultWriter io.Writer) io.Writer {
	if value := ctx.Value(stderrKey{}); value != nil {
		if typed, ok := value.(io.Writer); ok {
			return typed
		}
	}
	return defaultWriter
}

type stdinKey struct{}

// WithStdin sets the stdin reader on a given context.
func WithStdin(ctx context.Context, reader io.Reader) context.Context {
	return context.WithValue(ctx, stdinKey{}, reader)
}

// GetStdin gets a stdin reader from a context, or a default.
func GetStdin(ctx context.Context, defaultReader io.Reader) io.Reader {
	if value := ctx.Value(stdinKey{}); value != nil {
		if typed, ok := value.(io.Reader); ok {
			return typed
		}
	}
	return defaultReader
}

type verboseKey struct{}

// WithVerbose sets the verbose flag on a given context.
func WithVerbose(ctx context.Context, verbose bool) context.Context {
	return context.WithValue(ctx, verboseKey{}, verbose)
}

// IsVerbose gets the verbose flag on a given context.
func IsVerbose(ctx context.Context) bool {
	if value := ctx.Value(verboseKey{}); value != nil {
		if typed, ok := value.(bool); ok {
			return typed
		}
	}
	return false
}

type debugKey struct{}

// WithDebug sets the debug flag on a given context.
func WithDebug(ctx context.Context, debug bool) context.Context {
	return context.WithValue(ctx, debugKey{}, debug)
}

// IsDebug gets the debug flag on a given context.
func IsDebug(ctx context.Context) bool {
	if value := ctx.Value(debugKey{}); value != nil {
		if typed, ok := value.(bool); ok {
			return typed
		}
	}
	return false
}
