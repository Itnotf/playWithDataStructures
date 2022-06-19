package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	assert.Equal(t, 5, fib(5))
	assert.Equal(t, 1, fib(2))
}
