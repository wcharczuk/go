package cmd

// FlagStringsBinder is a flag that takes string values.
type FlagStringsBinder interface {
	Default() []string
	Value() *[]string
}

// FlagStringsP returns a new flag.
func FlagStringsP(value *[]string, name, short string, defaultValue []string, usage string) *FlagStringsVar {
	var fs FlagStringsVar
	fs.Field.Name = name
	fs.Field.Short = short
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagStringsVarP is a flag value with a shorthand.
type FlagStringsVarP struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default []string
		Value   *[]string
	}
}

// Name returns the flag name.
func (fs *FlagStringsVarP) Name() string { return fs.Field.Name }

// Short returns the flag short name.
func (fs *FlagStringsVarP) Short() string { return fs.Field.Short }

// Usage returns the flag usage.
func (fs *FlagStringsVarP) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagStringsVarP) Default() []string { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagStringsVarP) Value() *[]string { return fs.Field.Value }

// FlagStrings returns a new flag.
func FlagStrings(value *[]string, name string, defaultValue []string, usage string) *FlagStringsVar {
	var fs FlagStringsVar
	fs.Field.Name = name
	fs.Field.Usage = usage
	fs.Field.Default = defaultValue
	fs.Field.Value = value
	return &fs
}

// FlagStringsVar is a flag value.
type FlagStringsVar struct {
	Field struct {
		Name    string
		Short   string
		Usage   string
		Default []string
		Value   *[]string
	}
}

// Name returns the flag name.
func (fs *FlagStringsVar) Name() string { return fs.Field.Name }

// Usage returns the flag usage.
func (fs *FlagStringsVar) Usage() string { return fs.Field.Usage }

// Default returns the flag default.
func (fs *FlagStringsVar) Default() []string { return fs.Field.Default }

// Value returns the value reference.
func (fs *FlagStringsVar) Value() *[]string { return fs.Field.Value }
