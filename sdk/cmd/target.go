package cmd

import "context"

// Target is a type that implements the command interface.
type Target interface {
	Name() string
	Action(context.Context) error
}

// TargetArgs is a target that provides expected args.
type TargetArgs interface {
	Args() []string
}

// TargetArgsOptional is a target that provides if args are optional.
type TargetArgsOptional interface {
	ArgsOptional() bool
}

// TargetDescription is a target that provides a description.
type TargetDescription interface {
	Description() string
}

// TargetUsage is a target that provides usage.
type TargetUsage interface {
	Usage() string
}

// TargetAliases is a target that provides aliases.
type TargetAliases interface {
	Aliases() []string
}

// TargetFlags is a type that returns a set of flags to bind to.
type TargetFlags interface {
	Flags() []Flag
}

// TargetContext is a type that returns a base context.
type TargetContext interface {
	Context() context.Context
}
