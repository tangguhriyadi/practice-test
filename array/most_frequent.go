package array

func MostFrequent(arr []int) int {
	frequencyMap := make(map[int]int)

	for _, item := range arr {
		frequencyMap[item]++
	}

	maxFrequency := 0
	mostFrequentItem := 0

	for item, freq := range frequencyMap {
		if freq > maxFrequency {
			maxFrequency = freq
			mostFrequentItem = item
		}
	}

	return mostFrequentItem
}

func MaxNum(arr []int) int {
	max := 0

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}
