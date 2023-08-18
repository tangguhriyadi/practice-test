package practicetest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	result := IsValidParentheses("{{}}")
	result2 := IsValidParentheses("]]")

	assert.Equal(t, true, result)
	assert.Equal(t, false, result2)

}
