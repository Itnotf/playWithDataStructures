package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	assert.Equal(t, true, checkInclusion("adc", "dcda"))
	assert.Equal(t, true, checkInclusion("a", "ab"))
	assert.Equal(t, true, checkInclusion("ab", "eidbaooo"))
	assert.Equal(t, false, checkInclusion("ab", "eidboaoo"))
	assert.Equal(t, false, checkInclusion("abc", "ab"))
}
