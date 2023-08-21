package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	dictionary := []string{"i", "love", "coding", "and", "gaming"}
	case1 := "Ilovecoding"
	expect := "I love coding"

	result := WordBreak(case1, dictionary)

	assert.Equal(t, expect, result)

}
