package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charBook := map[byte]uint16{}

	for i := 0; i < len(s); i++ {
		charBook[s[i]]++
		charBook[t[i]]--
	}

	for _, count := range charBook {
		if count != 0 {
			return false
		}
	}

	return true
}
