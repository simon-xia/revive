package test

import (
	"testing"

	"github.com/simon-xia/revive/lint"
	"github.com/simon-xia/revive/rule"
)

func TestArgumentLimit(t *testing.T) {
	testRule(t, "argument-limit", &rule.ArgumentsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
