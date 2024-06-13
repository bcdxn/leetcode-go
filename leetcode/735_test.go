package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 10, -5},
			expected: []int{5, 10},
		},
		{
			input:    []int{8, -8},
			expected: []int{},
		},
		{
			input:    []int{10, 2, -5},
			expected: []int{10},
		},
		{
			input:    []int{-2, 1, -1, -2},
			expected: []int{-2, -2},
		},
		{
			input:    []int{1, -1, -2, -2},
			expected: []int{-2, -2},
		},
		{
			input:    []int{2, -1, -2, -2},
			expected: []int{-2},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, asteroidCollision(tc.input))
		})
	}
}
