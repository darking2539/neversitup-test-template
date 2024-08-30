package main

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	testCases := []TestCase{
		{
			CaseName: "case 1: one character Input",
			Input:    "a",
			Expected: []string{"a"},
		},
		{
			CaseName: "case 2: two characters Input",
			Input:    "ab",
			Expected: []string{"ab", "ba"},
		},
		{
			CaseName: "case 3: three characters Input",
			Input:    "abc",
			Expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			CaseName: "case 3,5: three characters Input",
			Input:    "cba",
			Expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			CaseName: "case 4: four characters Input",
			Input:    "aabb",
			Expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			CaseName: "case 5: same two characters",
			Input:    "aa",
			Expected: []string{"aa"},
		},
		{
			CaseName: "case 6: same three characters",
			Input:    "aaa",
			Expected: []string{"aaa"},
		},
		{	CaseName: "case 7: four characters Input",
			Input:    "abcd",
			Expected: []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
		},
	}

	for _, tc := range testCases {
        assert.Equal(t, tc.Expected, Permutations(tc.Input.(string)), tc.CaseName)
	}
}
