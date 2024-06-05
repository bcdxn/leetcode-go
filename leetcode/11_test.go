package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			input:    []int{1, 1},
			expected: 1,
		},
		{
			input:    []int{2, 3, 10, 5, 7, 8, 9},
			expected: 36,
		},
		{
			input:    []int{2, 3, 4, 5, 18, 17, 6},
			expected: 17,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxArea(tc.input))
		})
	}
}
