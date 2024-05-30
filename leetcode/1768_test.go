package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/bcdxn/leetcode-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMergeStringsAlternatively(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"abc", "pqr"},
			expected: "apbqcr",
		},
		{
			input:    []string{"ab", "pqrs"},
			expected: "apbqrs",
		},
		{
			input:    []string{"abcd", "pq"},
			expected: "apbqcd",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, leetcode.MergeAlternately(tc.input[0], tc.input[1]))
		})
	}
}
