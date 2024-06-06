package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOperations(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 2, 3, 4}, {5}},
			expected: 2,
		},
		{
			input:    [][]int{{3, 1, 3, 4, 3}, {6}},
			expected: 1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxOperations(tc.input[0], tc.input[1][0]))
		})
	}
}
