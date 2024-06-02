package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello",
			expected: "holle",
		},
		{
			input:    "leetcode",
			expected: "leotcede",
		},
		{
			input:    "a.",
			expected: "a.",
		},
		{
			input:    "aA",
			expected: "Aa",
		},
		{
			input:    "Marge, let's \"went.\" I await news telegram.",
			expected: "Marge, let's \"went.\" i awaIt news telegram.",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, reverseVowels(tc.input))
		})
	}
}
