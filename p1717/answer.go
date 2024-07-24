package p1717

func maximumGain(s string, x int, y int) int {

	if x > y {
		return _maximumGain(s, x, y, 0)
	} else {
		return _maximumGain(s, x, y, 1)
	}

}

func _maximumGain(s string, x int, y int, priorType int) int {
	if len(s) < 2 {
		return 0
	}

	blockByType := [][]byte{{'a', 'b'}, {'b', 'a'}}
	scoreByType := []int{x, y}

	var totalPoints int

	offset := 0

	badCount := 0
	for {

		if badCount == 2 {
			break
		}

		remainStr, gainPoints, ok, o1 := removeBlock(s, scoreByType[priorType], blockByType[priorType], offset)
		if ok {
			s = remainStr
			totalPoints += gainPoints
			badCount = 0
		} else {
			badCount++
		}

		if badCount == 2 {
			break
		}
		remainStr, gainPoints, ok, o2 := removeBlock(s, scoreByType[1-priorType], blockByType[1-priorType], 0)
		if ok {
			s = remainStr
			totalPoints += gainPoints
			badCount = 0
		} else {
			badCount++
		}

		if o2 == -1 {
			offset = o1
		} else if o1 >= o2 {
			offset = o2
		} else if o1 < o2 {
			offset = o1
		}
	}

	return totalPoints
}

func removeBlock(s string, score int, block []byte, offset int) (string, int, bool, int) {
	originalLen := len(s)
	totalScore := 0

	minChangedIdx := -1

	i := 0
	n := max(originalLen-1, 0)

	for {
		if i == n {
			break
		}

		if s[i] == block[0] && s[i+1] == block[1] {
			totalScore += score

			j := i + 1 // init
			for {
				if i-1 < 0 || j+1 > n {
					break
				}

				if s[i-1] == block[0] && s[j+1] == block[1] {
					totalScore += score
					i--
					j++
				} else {
					break
				}
			}

			if j+1 <= n {
				s = s[:i] + s[j+1:]
			} else {
				s = s[:i]
			}

			n = max(0, len(s)-1)
			i = max(i-1, 0)
			minChangedIdx = i

		} else {
			i++
		}
	}

	return s, totalScore, originalLen != len(s), minChangedIdx
}
