package DP

/*
The Tribonacci sequence Tn is defined as follows:
T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
Given n, return the value of Tn.
*/

func Tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	s := [3]int{0, 1, 1}
	for i := 3; i <= n; i++ {
		val := s[0] + s[1] + s[2]
		s[0] = s[1]
		s[1] = s[2]
		s[2] = val
	}
	return s[2]
}
