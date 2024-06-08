package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3},
			expected: -1,
		},
		{
			input:    []int{2, 1, -1},
			expected: 0,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, pivotIndex(tc.input))
		})
	}
}
