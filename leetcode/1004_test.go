package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, {2}},
			expected: 6,
		},
		{
			input:    [][]int{{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {3}},
			expected: 10,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, longestOnes(tc.input[0], tc.input[1][0]))
		})
	}
}
