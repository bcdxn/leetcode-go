package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			input:    []int{-4, -3, -2, -1, 4, 3, 2},
			expected: 0,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, largestAltitude(tc.input))
		})
	}
}

/* Info
------------------------------------------------------------------------------------------------- */

// # Description:
//
// There is a biker going on a road trip. The road trip consists of n + 1 points at different
// altitudes. The biker starts his trip on point 0 with altitude equal 0.
//
// You are given an integer array gain of length n where gain[i] is the net gain in altitude between
// points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.
//
// https://leetcode.com/problems/find-the-highest-altitude
//
// tags: [prefixsum, easy]
