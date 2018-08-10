package test

import (
	"testing"

	"github.com/simon-xia/revive/lint"
	"github.com/simon-xia/revive/rule"
)

func TestDisabledAnnotations(t *testing.T) {
	testRule(t, "disable-annotations", &rule.ExportedRule{}, &lint.RuleConfig{})
}
