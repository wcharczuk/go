package cmd

import (
	"context"
)

var (
	_ Target             = (*TargetBuilder)(nil)
	_ TargetArgs         = (*TargetBuilder)(nil)
	_ TargetArgsOptional = (*TargetBuilder)(nil)
	_ TargetDescription  = (*TargetBuilder)(nil)
	_ TargetUsage        = (*TargetBuilder)(nil)
	_ TargetFlags        = (*TargetBuilder)(nil)
	_ TargetAliases      = (*TargetBuilder)(nil)
)

// NewTarget returns a target from a name and an optional set of options.
func NewTarget(name string, opts ...TargetBuilderOption) *TargetBuilder {
	command := new(TargetBuilder)
	command.Field.Name = name
	for _, option := range opts {
		option(command)
	}
	return command
}

// TargetBuilderOption is a mutator for a target.
type TargetBuilderOption func(*TargetBuilder)

// OptTargetDescription sets the target description.
func OptTargetDescription(description string) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Description = description
	}
}

// OptTargetAliases adds to the target alias(es).
func OptTargetAliases(aliases ...string) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Aliases = append(c.Field.Aliases, aliases...)
	}
}

// OptTargetArgs sets the target expected args.
func OptTargetArgs(args ...string) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Args = args
	}
}

// OptTargetArgsOptional sets if the target expected args are optional.
func OptTargetArgsOptional(optional bool) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.ArgsOptional = optional
	}
}

// OptTargetUsage sets the target usage.
func OptTargetUsage(usage string) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Usage = usage
	}
}

// OptTargetSteps sets the target steps.
func OptTargetSteps(steps ...Action) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Steps = steps
	}
}

// OptTargetCleanupSteps sets the target steps.
func OptTargetCleanupSteps(steps ...Action) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.CleanupSteps = steps
	}
}

// OptTargetFlags sets the target flags.
func OptTargetFlags(flags ...Flag) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Flags = flags
	}
}

// OptTargetContext sets the target context.
func OptTargetContext(ctx context.Context) TargetBuilderOption {
	return func(c *TargetBuilder) {
		c.Field.Context = ctx
	}
}

// TargetBuilder is a generic target.
type TargetBuilder struct {
	Field struct {
		Name         string
		Context      context.Context
		Args         []string
		ArgsOptional bool
		Description  string
		Aliases      []string
		Usage        string
		Steps        []Action
		CleanupSteps []Action
		Flags        []Flag
	}
}

// Name returns the target name.
func (tb TargetBuilder) Name() string {
	return tb.Field.Name
}

// Aliases returns the target aliases.
func (tb TargetBuilder) Aliases() []string {
	return tb.Field.Aliases
}

// Args returns the target expected args.
func (tb TargetBuilder) Args() []string {
	return tb.Field.Args
}

// ArgsOptional returns if the target are optional.
func (tb TargetBuilder) ArgsOptional() bool {
	return tb.Field.ArgsOptional
}

// Description returns the target description.
func (tb TargetBuilder) Description() string {
	return tb.Field.Description
}

// Usage returns the target usage.
func (tb TargetBuilder) Usage() string {
	return tb.Field.Usage
}

// Flags returns the target flags.
func (tb TargetBuilder) Flags() []Flag {
	return tb.Field.Flags
}

// Action is the target body.
func (tb TargetBuilder) Action(ctx context.Context) (err error) {
	if len(tb.Field.CleanupSteps) > 0 {
		defer func() {
			var cleanupErr error
			for _, step := range tb.Field.CleanupSteps {
				if cleanupErr = step(ctx); cleanupErr != nil {
					Errorf(ctx, "%+v\n", cleanupErr)
				}
			}
		}()
	}
	for _, step := range tb.Field.Steps {
		if err = step(ctx); err != nil {
			return
		}
	}
	return
}
