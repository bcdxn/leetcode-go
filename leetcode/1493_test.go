package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			input:    []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			input:    []int{1, 1, 1},
			expected: 2,
		},
		{
			input:    []int{0, 0, 0},
			expected: 0,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, longestSubarray(tc.input))
		})
	}
}
