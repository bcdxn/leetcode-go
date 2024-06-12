package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualPairs(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			expected: 1,
		},
		{
			input:    [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			expected: 3,
		},
		{
			input:    [][]int{{11, 1}, {1, 11}},
			expected: 2,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, equalPairs(tc.input))
		})
	}
}
