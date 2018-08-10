package test

import (
	"testing"

	"github.com/simon-xia/revive/rule"
)

func TestUnusedParam(t *testing.T) {
	testRule(t, "unused-param", &rule.UnusedParamRule{})
}
