package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheOddInt(t *testing.T) {
	testCases := []TestCase{
		{
			CaseName: "which is odd.",
			Input:    []int{7},
			Expected: 7,
		},
		{
			CaseName: "which is odd.",
			Input:    []int{0},
			Expected: 0,
		},
		{
			CaseName: "which is odd.",
			Input:    []int{1, 1, 2},
			Expected: 2,
		},
		{
			CaseName: "which is odd.",
			Input:    []int{0, 1, 0, 1, 0},
			Expected: 0,
		},
		{
			CaseName: "should return 4, because it appears 1 time (which is odd).",
			Input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			Expected: 4,
		},
		{
			CaseName: "unExpected error",
			Input:    []int{1, 2},
			Expected: -1,
		},
		{
			CaseName: "unExpected error",
			Input:    []int{},
			Expected: -1,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.Expected, FindOddNumber(tc.Input.([]int)))
	}
}
