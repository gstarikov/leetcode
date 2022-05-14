package ms_backtraking

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Input: digits = ""
Output: []

Input: digits = "2"
Output: ["a","b","c"]

Constraints:
    0 <= digits.length <= 4
    digits[i] is a digit in the range ['2', '9'].

*/

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	var ret []string
	nextLevelStrings := letterCombinations(digits[1:])
	prefix := lettersAtDigit[digits[0]]
	if len(nextLevelStrings) > 0 {
		for _, pref := range prefix {
			for _, v := range nextLevelStrings {
				tmp := string(pref) + v
				ret = append(ret, tmp)
			}
		}
	} else {
		for _, pref := range prefix {
			ret = append(ret, string(pref))
		}
	}

	return ret
}

var lettersAtDigit = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}
