package ms_arrays_and_strings

import (
	"unicode"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
*/

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		switch {
		case !(unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))):
			i++
			continue
		case !(unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))):
			j--
			continue
		case unicode.ToUpper(rune(s[i])) != unicode.ToUpper(rune(s[j])):
			return false
		default:
			i++
			j--
		}
	}
	return true
}
