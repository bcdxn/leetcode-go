package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	r := Constructor()

	c := r.Ping(1)
	assert.Equal(t, 1, c)
	c = r.Ping(100)
	assert.Equal(t, 2, c)
	c = r.Ping(3001)
	assert.Equal(t, 3, c)
	c = r.Ping(3002)
	assert.Equal(t, 3, c)
}
