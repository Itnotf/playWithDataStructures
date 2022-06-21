package HJ5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_conv16To10(t *testing.T) {
	assert.Equal(t, 170, conv16To10("0xAA"))
}
