package p53

func maxSubArray(nums []int) int {

	cache := map[int]map[int]int{} // begin -> end -> sum

	// sliding window, size from 1 to len(nums)
	maximum := 0
	for size := 1; size <= len(nums); size++ {
		for i, j := 0, size-1; j < len(nums); i, j = i+1, j+1 {
			maximum = max(maximum, sumOfSubArray(nums, i, j, cache))
		}
	}

	return maximum

}

func sumOfSubArray(nums []int, i, j int, cache map[int]map[int]int) int {
	var sum int
	switch j - i {
	case 0:
		return nums[i]
	case 1:
		sum = nums[i] + nums[j]
	default:
		sum = cache[i][j-1] + nums[j]
	}

	m, ok := cache[i]
	if ok {
		m[j] = sum
	} else {
		cache[i] = map[int]int{j: sum}
	}
	return sum

}
