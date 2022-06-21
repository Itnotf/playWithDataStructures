package HJ2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getCountOfChar(t *testing.T) {
	assert.Equal(t, 2, getCountOfChar("ABCabc", "A"))
}
