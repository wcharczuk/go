package cmd

// Flag is a commanline flag.
type Flag interface {
	Name() string
	Usage() string
}

// FlagP is a commanline flag with a shorthand.
type FlagP interface {
	Flag
	Short() string
}
