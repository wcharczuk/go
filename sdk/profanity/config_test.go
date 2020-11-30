package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestConfig(t *testing.T) {
	its := assert.New(t)

	trueValue := true

	cfg := Config{}
	its.False(cfg.DebugOrDefault())
	cfg.Debug = &trueValue
	its.True(cfg.DebugOrDefault())

	its.False(cfg.VerboseOrDefault())
	cfg.Verbose = &trueValue
	its.True(cfg.VerboseOrDefault())

	its.False(cfg.FailFastOrDefault())
	cfg.FailFast = &trueValue
	its.True(cfg.FailFastOrDefault())

	its.Equal(DefaultRulesFile, cfg.RulesFileOrDefault())
	cfg.RulesFile = "foo"
	its.Equal("foo", cfg.RulesFileOrDefault())
}
