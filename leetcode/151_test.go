package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			input:    "a good   example",
			expected: "example good a",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, reverseWords(tc.input))
		})
	}
}
