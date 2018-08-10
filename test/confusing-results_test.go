package test

import (
	"testing"

	"github.com/simon-xia/revive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing-results", &rule.ConfusingResultsRule{})
}
