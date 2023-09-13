package array

import "sort"

func ArraySubset(arr []int) []int {
	totalSum := 0

	for _, num := range arr {
		totalSum += num
	}

	sort.Ints(arr)

	subsetA := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		currentNum := arr[i]
		subsetA = append(subsetA, currentNum)

		currentSum := 0

		for _, num := range subsetA {
			currentSum += num
		}

		currentResSum := totalSum - currentSum

		if currentSum > currentResSum {
			break
		}

	}

	return subsetA

}
