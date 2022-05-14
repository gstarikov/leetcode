package ms_arrays_and_strings

/*
Given a string s, return the longest palindromic substring in s.
*/

func LongestPalindrome(s string) string {
	switch len(s) {
	case 0, 1:
		return s
	}

	// run i := 0:len
	// on every elem run to left and right to check polindrome

	var p1, p2 int
	b := []byte(s)
	for i := range s {
		t1, t2 := checkPolindrome(b, i)
		if (t2 - t1) > (p2 - p1) {
			p1, p2 = t1, t2
		}
	}
	return s[p1 : p2+1]
}

func checkPolindrome(s []byte, i int) (p1, p2 int) {
	//check single root
	p1, p2 = checkPolindromeRoot(s, i, i)
	t1, t2 := p1, p2
	if i+1 < len(s) {
		t1, t2 = checkPolindromeRoot(s, i, i+1)
	}
	switch {
	case s[t1] != s[t2]: // double root is failed
		return p1, p2
	case (p2 - p1) > (t2 - t1): // single root is bigger
		return p1, p2
	default:
		return t1, t2
	}
}
func checkPolindromeRoot(s []byte, p1, p2 int) (int, int) {
	ln := len(s)
	t := p1
	for p1 >= 0 && p2 < ln {
		if s[p1] != s[p2] {
			break
		}
		p1--
		p2++
	}
	if t != p1 { // we are out of range so correct it back
		p1++
		p2--
	}
	return p1, p2
}
