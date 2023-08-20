package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostFrequent(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5}
	result := MostFrequent(arr)

	assert.Equal(t, 5, result)
}

func TestMaxNum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	result := MaxNum(arr)

	assert.Equal(t, 5, result)

}
