package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestConfig(t *testing.T) {
	its := assert.New(t)

	boolTrue := true

	cfg := Config{}
	its.False(cfg.DebugOrDefault())
	cfg.Debug = &boolTrue
	its.True(cfg.DebugOrDefault())

	its.False(cfg.VerboseOrDefault())
	cfg.Verbose = &boolTrue
	its.True(cfg.VerboseOrDefault())

	its.False(cfg.FailFastOrDefault())
	cfg.FailFast = &boolTrue
	its.True(cfg.FailFastOrDefault())

	its.Equal(DefaultRulesFile, cfg.RulesFileOrDefault())
	cfg.RulesFile = "foo"
	its.Equal("foo", cfg.RulesFileOrDefault())
}
