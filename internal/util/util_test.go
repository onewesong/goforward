package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRangeNumbers(t *testing.T) {
	assert := assert.New(t)
	numbers, err := ParseRangeNumbers("2-5")
	if assert.NoError(err) {
		assert.Equal([]int64{2, 3, 4, 5}, numbers)
	}

	numbers, err = ParseRangeNumbers("1")
	if assert.NoError(err) {
		assert.Equal([]int64{1}, numbers)
	}

	numbers, err = ParseRangeNumbers("3-5,8")
	if assert.NoError(err) {
		assert.Equal([]int64{3, 4, 5, 8}, numbers)
	}

	numbers, err = ParseRangeNumbers(" 3-5,8, 10-12 ")
	if assert.NoError(err) {
		assert.Equal([]int64{3, 4, 5, 8, 10, 11, 12}, numbers)
	}

	_, err = ParseRangeNumbers("3-a")
	assert.Error(err)
}
