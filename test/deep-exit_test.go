package test

import (
	"testing"

	"github.com/simon-xia/revive/rule"
)

func TestDeepExit(t *testing.T) {
	testRule(t, "deep-exit", &rule.DeepExitRule{})
}
