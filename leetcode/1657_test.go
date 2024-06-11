package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	cases := []struct {
		input    []string
		expected bool
	}{
		{
			input:    []string{"abc", "bca"},
			expected: true,
		},
		{
			input:    []string{"a", "aa"},
			expected: false,
		},
		{
			input:    []string{"cabbba", "abbccc"},
			expected: true,
		},
		{
			input:    []string{"cabbba", "aabbss"},
			expected: false,
		},
		{
			input:    []string{"abbbzcf", "babzzcz"},
			expected: false,
		},
		{
			input:    []string{"uau", "ssx"},
			expected: false,
		},
		{
			input:    []string{"aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff"},
			expected: false,
		},
		{
			input:    []string{"abbzccc", "babzzcz"},
			expected: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, closeStrings(tc.input[0], tc.input[1]))
		})
	}
}
