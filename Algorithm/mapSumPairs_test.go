package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSum_sum(t1 *testing.T) {
	mapSum := MapSumConstructor()
	mapSum.Insert("apple", 3)
	assert.Equal(t1, 3, mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	assert.Equal(t1, 5, mapSum.Sum("ap"))
	assert.Equal(t1, 5, mapSum.Sum("abbb"))
}
