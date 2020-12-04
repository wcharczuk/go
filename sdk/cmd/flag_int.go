package cmd

// FlagIntBinder is a flag that takes bool values.
type FlagIntBinder interface {
	Default() int
	Value() *int
}

// FlagIntP returns a new flag.
func FlagIntP(value *int, name, short string, defaultValue int, usage string) *FlagIntVarP {
	var fs FlagIntVarP
	fs.Field.Name = name
	fs.Field.Short = short
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagInt returns a new flag.
func FlagInt(value *int, name string, defaultValue int, usage string) *FlagIntVar {
	var fs FlagIntVar
	fs.Field.Name = name
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagIntVarP is a flag value.
type FlagIntVarP struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default int
		Value   *int
	}
}

// Name returns the flag name.
func (fs *FlagIntVarP) Name() string { return fs.Field.Name }

// Short returns the flag short name.
func (fs *FlagIntVarP) Short() string { return fs.Field.Short }

// Usage returns the flag usage.
func (fs *FlagIntVarP) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagIntVarP) Default() int { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagIntVarP) Value() *int { return fs.Field.Value }

// FlagIntVar is a flag value.
type FlagIntVar struct {
	Field struct {
		Name    string
		Usage   string
		Default int
		Value   *int
	}
}

// Name returns the flag name.
func (fs *FlagIntVar) Name() string { return fs.Field.Name }

// Usage returns the flag usage.
func (fs *FlagIntVar) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagIntVar) Default() int { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagIntVar) Value() *int { return fs.Field.Value }
