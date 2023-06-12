package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type errorTestCases struct {
		description      string
		input            string
		expectedResponse bool
	}

	for _, scenario := range []errorTestCases{
		{
			description:      "()",
			input:            "()",
			expectedResponse: true,
		},
		{
			description:      "()[]{}",
			input:            "()[]{}",
			expectedResponse: true,
		},
		{
			description:      "(]",
			input:            "(]",
			expectedResponse: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			resp := isValid(scenario.input)
			assert.Equal(t, scenario.expectedResponse, resp)
		})
	}
}
