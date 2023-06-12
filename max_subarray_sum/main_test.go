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
		input            []int
		expectedResponse int
	}

	for _, scenario := range []errorTestCases{
		{
			description:      "{-2, -3, 4, -1, -2, 1, 5, -3}",
			input:            []int{-2, -3, 4, -1, -2, 1, 5, -3},
			expectedResponse: 7,
		},
		{
			description:      "{-2, -3, 4, -1, -2, 1, 5, -3, 4, -1, -2, 1, 10}",
			input:            []int{-2, -3, 4, -1, -2, 1, 5, -3, 4, -1, -2, 1, 10},
			expectedResponse: 16,
		},
		{
			description:      "{5, 4, -1, 7, 8}",
			input:            []int{5, 4, -1, 7, 8},
			expectedResponse: 23,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			resp := maxSubArray(scenario.input)
			assert.Equal(t, scenario.expectedResponse, resp)
		})
	}
}
