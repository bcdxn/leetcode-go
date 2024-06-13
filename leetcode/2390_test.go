package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "leet**cod*e",
			expected: "lecoe",
		},
		{
			input:    "erase*****",
			expected: "",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, removeStars(tc.input))
		})
	}
}
