package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmilyFace(t *testing.T) {

	testCases := []TestCase{
		{
			CaseName: "case 1",
			Input:    []string{":)", ";(", ";}", ":-D"},
			Expected: 2,
		},
		{
			CaseName: "case 2",
			Input:    []string{";D", ":-(", ":-)", ";~)"},
			Expected: 3,
		},
		{
			CaseName: "case 3",
			Input:    []string{";]", ":[", ";*", ":$", ";-D"},
			Expected: 1,
		},
		{
			CaseName: "case 4",
			Input:    []string{":[", ";(", ";}", ":-|"},
			Expected: 0,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Expected, CountSmilyFace(tc.Input.([]string)))
	}
}