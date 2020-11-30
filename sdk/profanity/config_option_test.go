package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func Test_Config_Options(t *testing.T) {
	its := assert.New(t)

	cfg := &Config{}

	its.False(cfg.VerboseOrDefault())
	OptVerbose(true)(cfg)
	its.True(cfg.VerboseOrDefault())

	its.False(cfg.DebugOrDefault())
	OptDebug(true)(cfg)
	its.True(cfg.DebugOrDefault())

	its.False(cfg.FailFastOrDefault())
	OptFailFast(true)(cfg)
	its.True(cfg.FailFastOrDefault())

	its.Empty(cfg.Path)
	OptPath("../foo")(cfg)
	its.Equal("../foo", cfg.Path)

	its.Equal(DefaultRulesFile, cfg.RulesFileOrDefault())
	OptRulesFile("my_rules.yml")(cfg)

	its.Empty(cfg.Files.Include)
	OptFilesInclude("foo", "bar", "baz")(cfg)
	its.Equal([]string{"foo", "bar", "baz"}, cfg.Files.Include)

	its.Empty(cfg.Files.Exclude)
	OptFilesExclude("foo", "bar", "baz")(cfg)
	its.Equal([]string{"foo", "bar", "baz"}, cfg.Files.Exclude)
}
