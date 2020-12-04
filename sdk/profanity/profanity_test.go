package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func Test_Profanity_ReadRuleSpecsFile(t *testing.T) {
	its := assert.New(t)

	profanity := &Profanity{}

	rules, err := profanity.ReadRuleSpecsFile("testdata/rules.yml")
	its.Nil(err)
	its.NotEmpty(rules)
}
