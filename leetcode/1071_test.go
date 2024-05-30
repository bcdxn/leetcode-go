package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"ABCABC", "ABC"},
			expected: "ABC",
		},
		{
			input:    []string{"ABABAB", "ABAB"},
			expected: "AB",
		},
		{
			input:    []string{"LEET", "CODE"},
			expected: "",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, gcdOfStrings(tc.input[0], tc.input[1]))
		})
	}
}
