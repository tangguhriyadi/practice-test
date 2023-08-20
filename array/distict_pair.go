package array

import "fmt"

func CountDistinctPairsWithSum(arr []int, target int) int {
	pairCount := 0

	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			count := fmt.Sprintf("counting %d + %d", arr[i], arr[j])
			fmt.Println(count)
			if arr[j]+arr[i] == target {
				pairCount++
			}
		}
	}

	return pairCount

}

// [1,2,3,4,5] 6
