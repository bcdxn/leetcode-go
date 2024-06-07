package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	cases := []struct {
		input struct {
			s string
			k int
		}
		expected int
	}{
		{
			input: struct {
				s string
				k int
			}{s: "abciiidef", k: 3},
			expected: 3,
		},
		{
			input: struct {
				s string
				k int
			}{s: "aeiou", k: 2},
			expected: 2,
		},
		{
			input: struct {
				s string
				k int
			}{s: "leetcode", k: 3},
			expected: 2,
		},
		{
			input: struct {
				s string
				k int
			}{s: "ibpbhixfiouhdljnjfflpapptrxgcomvnb", k: 33},
			expected: 7,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxVowels(tc.input.s, tc.input.k))
		})
	}
}
