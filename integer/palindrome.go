package integer

import "math"

func isPalindrome(number int) bool {
	if number < 0 {
		return false
	}

	temp := number
	var reversed int

	for temp > 0 {
		lastDigit := temp % 10
		reversed = (reversed * 10) + lastDigit
		temp = int(math.Floor(float64(temp) / 10))
	}

	return reversed == number
}
