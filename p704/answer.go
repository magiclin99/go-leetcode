package p704

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	last := len(nums) - 1
	i := len(nums) / 2
	for {

		if nums[i] == target {
			return i
		} else if target < nums[i] {

			if i == 0 {
				return -1
			}

			last = i - 1
			i = last / 2

		} else {

			if i == last {
				return -1
			}

			sum := i + last
			i = sum / 2
			if sum%2 != 0 {
				i++
			}
		}

	}
}
