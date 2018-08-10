package test

import (
	"testing"

	"github.com/simon-xia/revive/rule"
)

func TestFlagParam(t *testing.T) {
	testRule(t, "flag-param", &rule.FlagParamRule{})
}
