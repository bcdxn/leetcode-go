package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			input:    "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, decodeString(tc.input))
		})
	}
}
