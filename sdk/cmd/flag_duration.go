package cmd

import "time"

// FlagDurationBinder is a flag that takes time.Duration values.
type FlagDurationBinder interface {
	Default() time.Duration
	Value() *time.Duration
}

// FlagDurationP returns a new flag.
func FlagDurationP(value *time.Duration, name, short string, defaultValue time.Duration, usage string) *FlagDurationVarP {
	var fs FlagDurationVarP
	fs.Field.Name = name
	fs.Field.Short = short
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagDuration returns a new flag.
func FlagDuration(value *time.Duration, name string, defaultValue time.Duration, usage string) *FlagDurationVar {
	var fs FlagDurationVar
	fs.Field.Name = name
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagDurationVarP is a flag value.
type FlagDurationVarP struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default time.Duration
		Value   *time.Duration
	}
}

// Name returns the flag name.
func (fs *FlagDurationVarP) Name() string { return fs.Field.Name }

// Short returns the flag short name.
func (fs *FlagDurationVarP) Short() string { return fs.Field.Short }

// Usage returns the flag usage.
func (fs *FlagDurationVarP) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagDurationVarP) Default() time.Duration { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagDurationVarP) Value() *time.Duration { return fs.Field.Value }

// FlagDurationVar is a flag value.
type FlagDurationVar struct {
	Field struct {
		Name    string
		Usage   string
		Default time.Duration
		Value   *time.Duration
	}
}

// Name returns the flag name.
func (fs *FlagDurationVar) Name() string { return fs.Field.Name }

// Usage returns the flag usage.
func (fs *FlagDurationVar) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagDurationVar) Default() time.Duration { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagDurationVar) Value() *time.Duration { return fs.Field.Value }
