package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlantFlowers(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected bool
	}{
		{
			input:    [][]int{{1, 0, 0, 0, 1}, {1}},
			expected: true,
		},
		{
			input:    [][]int{{1, 0, 0, 0, 1}, {2}},
			expected: false,
		},
		{
			input:    [][]int{{1, 0, 0, 0, 0, 1}, {2}},
			expected: false,
		},
		{
			input:    [][]int{{1, 0, 1, 0, 1, 0, 1}, {0}},
			expected: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, canPlaceFlowers(tc.input[0], tc.input[1][0]))
		})
	}
}
