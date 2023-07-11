package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	type errorTestCases struct {
		description      string
		inputString      string
		inputPattern     string
		expectedResponse bool
	}

	for _, scenario := range []errorTestCases{
		{
			description:      "s = \"aa\", p = \"a\"",
			inputString:      "aa",
			inputPattern:     "a",
			expectedResponse: false,
		},
		{
			description:      "s = \"aa\", p = \"a*\"",
			inputString:      "aa",
			inputPattern:     "a*",
			expectedResponse: true,
		},
		{
			description:      "s = \"ab\", p = \".*\"",
			inputString:      "ab",
			inputPattern:     ".*",
			expectedResponse: true,
		},
		{
			description:      "s = \"abbccc\", p = \"a*bc*\"",
			inputString:      "abbccc",
			inputPattern:     "a*bc*",
			expectedResponse: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			resp := isMatch(scenario.inputString, scenario.inputPattern)
			assert.Equal(t, scenario.expectedResponse, resp)
		})
	}
}
