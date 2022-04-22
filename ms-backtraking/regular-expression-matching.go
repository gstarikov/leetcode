package ms_backtraking

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
    '.' Matches any single character.
    '*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".

Constraints:

    1 <= s.length <= 20
    1 <= p.length <= 30
    s contains only lowercase English letters.
    p contains only lowercase English letters, '.', and '*'.
    It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

func isMatchDP(s string, p string) bool {
	//try DP to solve a task
	maxS := len(s) + 1
	maxP := len(p) + 1

	memo := make([]state, maxS*maxP)
	getDP := func(i, j int) state {
		idx := i*maxP + j
		return memo[idx]
	}
	setDP := func(i, j int, s bool) {
		idx := i*maxP + j
		switch s {
		case true:
			memo[idx] = stateTrue
		case false:
			memo[idx] = stateFalse
		default:
			panic("impossible case")
		}
	}
	var doDP func(i, j int, s, p *string) bool
	doDP = func(i, j int, s, p *string) bool { //pointer to string to avoid strings copy
		//check dp cache
		st := getDP(i, j)
		switch st {
		case stateEmpty:
		//do pattern check
		case stateTrue:
			return true
		case stateFalse:
			return false
		default:
			panic("impossible case")
		}

		var resp bool
		//do pattern check
		switch {
		case j == len(*p):
			return i == len(*s)
		case i < len(*s):
			match := (*s)[i] == (*p)[j] || (*p)[j] == '.'
			star := j+1 < len(*p) && (*p)[j+1] == '*'

			switch {
			case star && match:
				patSkip := doDP(i, j+2, s, p)   //check hypothesis of pattern skip
				patNoSkip := doDP(i+1, j, s, p) //continue pattern check
				resp = patSkip || patNoSkip

			case star && !match:
				resp = doDP(i, j+2, s, p) //skip wildcard
			case !star && match:
				resp = doDP(i+1, j+1, s, p)
			case !star && !match:
				resp = false
			}

		case i == len(*s): // pattern isnt ended, but string is ended
			star := j+1 < len(*p) && (*p)[j+1] == '*'
			if star {
				resp = doDP(i, j+2, s, p)
			} else {
				resp = false
			}

		default:
			panic("impossible case")
		}
		setDP(i, j, resp)
		return resp
	}
	return doDP(0, 0, &s, &p)
}

type state int8

const (
	stateEmpty = iota
	stateTrue  = iota
	stateFalse = iota
)

func isMatch(s string, p string) bool {
	// scan pattern string
	// in case of .*
	// in case of any character plus star
	// in case of .
	// in all other cases (any other character)

	pPos := 0
	sPos := 0
	starPos := -1
	var prevStarChar byte

	for sPos < len(s) {
		if pPos < len(p) {
			switch {
			case len(p)-pPos > 1 && p[pPos+1] == '*': //check ?*
				prevStarChar = p[pPos]
				starPos = pPos + 2 //set to after star position
				pPos = pPos + 2
			case p[pPos] == '.':
				sPos++ //go to next
				pPos++
			//case p[pPos] == '*': //error case. star may be as ordinary sumbol
			//	panic("error...")
			case p[pPos] == s[sPos]: //pattern match. go next
				pPos++
				sPos++
			case p[pPos] != s[sPos]: //last try to  expand last star
				switch {
				case starPos == -1:
					return false
				case prevStarChar == '.', prevStarChar == s[sPos]:
					sPos++
					pPos = starPos
				default:
					return false
				}
			default:
				panic("impossible case")
			}
		} else { //check tail cases
			switch {
			case starPos == -1:
				return false
			case prevStarChar == '.', prevStarChar == s[sPos]:
				sPos++
				pPos = starPos
			default:
				return false
			}
		}
	}

	//check tail cases. it must be star pattern till the end
	switch {
	case sPos == len(s) && pPos == len(p):
		return true
	case pPos == len(p) && starPos != -1:
		//must be prevStarChar till the end
		for ; sPos < len(s); sPos++ {
			switch prevStarChar {
			case '.', s[sPos]: //do nothing
			default:
				return false
			}
		}
		return true
	default:
		return false
	}
}
