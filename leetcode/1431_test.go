package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKidsWithCandies(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []bool
	}{
		{
			input:    [][]int{{2, 3, 5, 1, 3}, {3}},
			expected: []bool{true, true, true, false, true},
		},
		{
			input:    [][]int{{4, 2, 1, 1, 2}, {1}},
			expected: []bool{true, false, false, false, false},
		},
		{
			input:    [][]int{{12, 1, 12}, {10}},
			expected: []bool{true, false, true},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, kidsWithCandies(tc.input[0], tc.input[1][0]))
		})
	}
}
