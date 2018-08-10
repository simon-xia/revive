package test

import (
	"testing"

	"github.com/simon-xia/revive/lint"
	"github.com/simon-xia/revive/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
