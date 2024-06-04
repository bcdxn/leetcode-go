package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			input:    []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, productExceptSelf(tc.input))
		})
	}
}
