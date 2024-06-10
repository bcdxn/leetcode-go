package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			input:    []int{1, 2},
			expected: false,
		},
		{
			input:    []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, uniqueOccurrences(tc.input))
		})
	}
}
