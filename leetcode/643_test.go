package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected float64
	}{
		{
			input:    [][]int{{1, 12, -5, -6, 50, 3}, {4}},
			expected: 12.75000,
		},
		{
			input:    [][]int{{5}, {1}},
			expected: 5.00000,
		},
		{
			input:    [][]int{{0, 4, 0, 3, 2}, {1}},
			expected: 4.00000,
		},
		{
			input:    [][]int{{4, 2, 1, 3, 3}, {2}},
			expected: 3.00000,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findMaxAverage(tc.input[0], tc.input[1][0]))
		})
	}
}
