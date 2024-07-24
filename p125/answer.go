package p125

func isPalindrome(s string) bool {

	var c1, c2 byte

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		for ; i < j && !isAlphabetNumeric(s[i]); i++ {
		}

		for ; i < j && !isAlphabetNumeric(s[j]); j-- {
		}

		c1 = s[i]
		c2 = s[j]

		if c1 >= 'a' && s[i] <= 'z' {
			c1 = c1 - 32
		}

		if c2 >= 'a' && s[j] <= 'z' {
			c2 = c2 - 32
		}

		if c1-c2 == 0 {
			continue
		} else {
			return false
		}
	}

	return true

}

func isAlphabetNumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
