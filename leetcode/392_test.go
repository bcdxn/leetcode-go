package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		input    []string
		expected bool
	}{
		{
			input:    []string{"abc", "ahbgdc"},
			expected: true,
		},
		{
			input:    []string{"axc", "ahbgdc"},
			expected: false,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isSubsequence(tc.input[0], tc.input[1]))
		})
	}
}
