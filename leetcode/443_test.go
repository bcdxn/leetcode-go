package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	cases := []struct {
		input    []byte
		expected int
	}{
		{
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			input:    []byte{'a'},
			expected: 1,
		},
		{
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, compress(tc.input))
		})
	}
}
