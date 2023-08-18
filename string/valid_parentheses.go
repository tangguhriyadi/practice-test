package string

import "fmt"

func IsValidParentheses(str string) bool {
	stack := make([]rune, 0)

	openingBracket := []rune{'{', '[', '('}
	closingBracket := []rune{'}', ']', ')'}

	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range str {
		if contains(openingBracket, char) {
			stack = append(stack, char)
		} else if contains(closingBracket, char) {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if bracketPairs[char] != top {
				return false
			}

		}
	}

	return len(stack) == 0

}

func contains(arr []rune, target rune) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}

	return false
}

func main() {
	result := IsValidParentheses("{{}}")

	fmt.Println(result)
}
