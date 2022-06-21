package NowCoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getLengthOfLastWord(t *testing.T) {
	assert.Equal(t, 8, getLengthOfLastWord("hello nowcoder"))
}
