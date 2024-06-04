package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{2, 1},
			expected: []int{2, 1},
		},
		{
			input:    []int{1, 0, 1},
			expected: []int{1, 1, 0},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			moveZeroes(tc.input)
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}
