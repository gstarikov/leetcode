package ms_arrays_and_strings

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

*/

func ValidParentheses(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0, len(s))
	for _, v := range s {
		switch v {
		case '(', '{', '[': //open
			stack = append(stack, byte(v))
		case ')', '}', ']': //close
			//check
			last := len(stack) - 1
			if last < 0 {
				return false
			}
			e := m[stack[last]]
			if e != byte(v) {
				return false
			}
			//remove
			stack = stack[:last]
		}
	}
	return len(stack) == 0
}
