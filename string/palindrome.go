package string

import (
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
	reg := regexp.MustCompile("[^a-zA-Z]")
	cleanedString := reg.ReplaceAllString(str, "")
	cleanedString = strings.ToLower(cleanedString)

	return cleanedString == ReverseString(cleanedString)
}

func ReverseString(str string) interface{} {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
