package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Args struct {
	src []int64
}

func TestGetMaxSubarray(t *testing.T) {
	type errorTestCases struct {
		description      string
		input            string
		expectedResponse string
	}

	for _, scenario := range []errorTestCases{
		{
			description:      "babad",
			input:            "babad",
			expectedResponse: "bab",
		},
		{
			description:      "cbbd",
			input:            "cbbd",
			expectedResponse: "bb",
		},
		{
			description:      "bb",
			input:            "bb",
			expectedResponse: "bb",
		},
		{
			description:      "tattarrattat",
			input:            "tattarrattat",
			expectedResponse: "tattarrattat",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			resp := longestPalindrome(scenario.input)
			assert.Equal(t, scenario.expectedResponse, resp)
		})
	}
}
