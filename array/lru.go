package array

func LRU(arr []int, count int) []int {
	temp := arr

	for i := 0; i < count; i++ {
		last := temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		temp = append([]int{last}, temp...)
	}

	return temp
}
