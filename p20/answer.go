package p20

/*
*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

type Stack []byte

func (it *Stack) Push(newChar byte) {
	*it = append(*it, newChar)
}

func (it *Stack) Pop() byte {
	if len(*it) == 0 {
		return 0
	}
	last := (*it)[len(*it)-1]
	*it = (*it)[:len(*it)-1]
	return last
}

func isValid(s string) bool {

	stack := Stack{}

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case '(', '[', '{':
			stack.Push(s[i])
		default:
			last := stack.Pop()
			if last == 0 {
				return false
			}
			if s[i] == ')' && last != '(' {
				return false
			}
			if s[i] == ']' && last != '[' {
				return false
			}
			if s[i] == '}' && last != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
