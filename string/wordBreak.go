package string

import (
	"fmt"
	"strings"
)

func WordBreak(str string, dictionary []string) string {

	word := ""

	result := ""

	for _, char := range str {
		word += string(char)
		if IsContaining(dictionary, word) {
			result += fmt.Sprintf("%s ", word)
			word = ""
		}
	}

	return strings.TrimSpace(result)
}

func IsContaining(dictionary []string, target string) bool {

	for _, char := range dictionary {
		if strings.EqualFold(char, target) {
			return true
		}
	}

	return false
}
