package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySubset(t *testing.T) {
	caseOne := []int{3, 7, 5, 6, 2}
	expectOne := []int{7, 6}

	result := ArraySubset(caseOne)

	assert.Equal(t, expectOne, result)
}
