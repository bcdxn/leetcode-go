package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeFindDifference(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 2, 3}, {2, 4, 6}},
			expected: [][]int{{1, 3}, {4, 6}},
		},
		{
			input:    [][]int{{1, 2, 3, 3}, {1, 1, 2, 2}},
			expected: [][]int{{3}, {}},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			actual := findDifference(tc.input[0], tc.input[1])
			assert.Equal(t, tc.expected[0], actual[0])
			assert.Equal(t, tc.expected[1], actual[1])
		})
	}
}
