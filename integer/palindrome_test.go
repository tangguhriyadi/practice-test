package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	result := isPalindrome(1001)
	result2 := isPalindrome(100)

	assert.Equal(t, true, result)
	assert.Equal(t, false, result2)
}
