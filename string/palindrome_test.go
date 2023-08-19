package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	result := isPalindrome("katak")
	result2 := isPalindrome("katak2!")
	result3 := isPalindrome("kataka2!")

	assert.Equal(t, true, result)
	assert.Equal(t, true, result2)
	assert.Equal(t, false, result3)

}
