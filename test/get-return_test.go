package test

import (
	"testing"

	"github.com/simon-xia/revive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
