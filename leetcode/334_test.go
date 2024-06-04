package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			input:    []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, increasingTriplet(tc.input))
		})
	}
}
