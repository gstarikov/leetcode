package backtraking

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
    '?' Matches any single character.
    '*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.

Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

Constraints:
    0 <= s.length, p.length <= 2000
    s contains only lowercase English letters.
    p contains only lowercase English letters, '?' or '*'.

*/

func WildcardMatching(s string, p string) bool {
	sIdx, pIdx := 0, 0
	starIdx := -1
	sTmpIdx := -1
	for sIdx < len(s) {
		switch {
		case pIdx < len(p) && (p[pIdx] == s[sIdx] || p[pIdx] == '?'): // go next
			sIdx++
			pIdx++
		case pIdx < len(p) && p[pIdx] == '*': //found star
			starIdx = pIdx
			pIdx++
			sTmpIdx = sIdx // check no one symbol match
		case starIdx == -1:
			return false
		default: //symbols dont match, but before we meet star
			pIdx = starIdx + 1 //reset pattern pointer
			sTmpIdx++          // one more symbol after star
			sIdx = sTmpIdx

		}
	}
	for ; pIdx < len(p) && p[pIdx] == '*'; pIdx++ {
	}
	return pIdx == len(p)
}
