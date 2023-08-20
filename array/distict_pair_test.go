package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	caseOne := []int{1, 2, 3, 4, 5}
	expectOne := 2
	result1 := CountDistinctPairsWithSum(caseOne, 6)

	assert.Equal(t, expectOne, result1)
}
