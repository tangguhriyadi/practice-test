package array

func CountDistinctPairsWithSum(arr []int, target int) int {
	pairCount := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]+arr[i] == target {
				pairCount++
			}
		}
	}

	// for i, j := 0, 1; i < len(arr) && j < len(arr); i, j = i+1, j+1 {
	// 	if arr[j]+arr[i] == target {
	// 		pairCount++
	// 	}
	// }

	return pairCount

}

// [1,2,3,4,5] 6
