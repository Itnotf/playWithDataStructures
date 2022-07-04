package Leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nthPersonGetsNthSeat(t *testing.T) {
	assert.Equal(t, 1.0, nthPersonGetsNthSeat(1))
	assert.Equal(t, 0.5, nthPersonGetsNthSeat(2))
}
