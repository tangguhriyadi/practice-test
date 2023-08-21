package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	count := 2

	result := LRU(arr, count)
	expect := []int{4, 5, 1, 2, 3}

	assert.Equal(t, expect, result)
}
