package cmd

// FlagBoolBinder is a flag that takes bool values.
type FlagBoolBinder interface {
	Default() bool
	Value() *bool
}

// FlagBoolP returns a new flag.
func FlagBoolP(value *bool, name, short string, defaultValue bool, usage string) *FlagBoolVarP {
	var fs FlagBoolVarP
	fs.Field.Name = name
	fs.Field.Short = short
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagBool returns a new flag.
func FlagBool(value *bool, name string, defaultValue bool, usage string) *FlagBoolVar {
	var fs FlagBoolVar
	fs.Field.Name = name
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagBoolVarP is a flag value.
type FlagBoolVarP struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default bool
		Value   *bool
	}
}

// Name returns the flag name.
func (fs *FlagBoolVarP) Name() string { return fs.Field.Name }

// Short returns the flag short name.
func (fs *FlagBoolVarP) Short() string { return fs.Field.Short }

// Usage returns the flag usage.
func (fs *FlagBoolVarP) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagBoolVarP) Default() bool { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagBoolVarP) Value() *bool { return fs.Field.Value }

// FlagBoolVar is a flag value.
type FlagBoolVar struct {
	Field struct {
		Name    string
		Usage   string
		Default bool
		Value   *bool
	}
}

// Name returns the flag name.
func (fs *FlagBoolVar) Name() string { return fs.Field.Name }

// Usage returns the flag usage.
func (fs *FlagBoolVar) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagBoolVar) Default() bool { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagBoolVar) Value() *bool { return fs.Field.Value }
