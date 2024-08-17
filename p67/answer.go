package p67

func addBinary(a string, b string) string {
	symbol := "01"

	bitTable := map[byte]int{
		symbol[0]: 0,
		symbol[1]: 1,
	}

	carry := false
	output := ""
	var tmp int
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {

		tmp = 0
		if i >= 0 {
			tmp += bitTable[a[i]]
		}
		if j >= 0 {
			tmp += bitTable[b[j]]
		}

		if carry {
			tmp += 1
		}

		switch tmp {
		case 0:
			output = "0" + output
			carry = false
		case 1:
			output = "1" + output
			carry = false
		case 2:
			output = "0" + output
			carry = true
		case 3:
			output = "1" + output
			carry = true
		}
	}

	if carry {
		output = "1" + output
	}

	return output
}
