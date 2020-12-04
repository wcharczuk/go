package cmd

// FlagStringBinder is a flag that takes string values.
type FlagStringBinder interface {
	Default() string
	Value() *string
}

// FlagStringP returns a new flag.
func FlagStringP(value *string, name, short, defaultValue, usage string) *FlagStringVarP {
	var fs FlagStringVarP
	fs.Field.Name = name
	fs.Field.Short = short
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagString returns a new flag.
func FlagString(value *string, name, defaultValue, usage string) *FlagStringVar {
	var fs FlagStringVar
	fs.Field.Name = name
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagStringVarP is a flag value with a shorthand.
type FlagStringVarP struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default string
		Value   *string
	}
}

// Name returns the flag name.
func (fs *FlagStringVarP) Name() string { return fs.Field.Name }

// Short returns the flag short name.
func (fs *FlagStringVarP) Short() string { return fs.Field.Short }

// Usage returns the flag usage.
func (fs *FlagStringVarP) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagStringVarP) Default() string { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagStringVarP) Value() *string { return fs.Field.Value }

// FlagStringVar is a flag value.
type FlagStringVar struct {
	Field struct {
		Name    string
		Usage   string
		Default string
		Value   *string
	}
}

// Name returns the flag name.
func (fs *FlagStringVar) Name() string { return fs.Field.Name }

// Usage returns the flag usage.
func (fs *FlagStringVar) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagStringVar) Default() string { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagStringVar) Value() *string { return fs.Field.Value }
