package p57

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	var output [][]int

	head, tail := 0, 0
	mode := 0 // 0, 1, 2
	var itv []int

	for i := 0; i < len(intervals); i++ {

		itv = intervals[i]

		switch mode {

		case 0: // find head
			if newInterval[0] <= itv[0] {
				head = newInterval[0]
				mode = 1
				i-- // let the loop check same itv again in mode 2
			} else if itv[0] < newInterval[0] && newInterval[0] <= itv[1] {
				head = itv[0]
				mode = 1
				i-- // let the loop check same itv again in mode 2
			} else {
				output = append(output, itv)
			}
		case 1: // find tail
			if newInterval[1] < itv[0] {
				tail = newInterval[1]
				output = append(output, []int{head, tail})
				output = append(output, intervals[i:]...)
				return output
			} else if newInterval[1] <= itv[1] {
				tail = itv[1]
				mode = 2
			} else if i == len(intervals)-1 {
				tail = newInterval[1]
				mode = 2
			}

			if mode == 2 {
				output = append(output, []int{head, tail})
				output = append(output, intervals[i+1:]...)
				return output
			} // end
		}
	}

	if mode == 0 {
		output = append(output, newInterval)
		return output
	}

	return output
}
