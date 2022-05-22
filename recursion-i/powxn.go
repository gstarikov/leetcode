package recursion_i

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Constraints:
    -100.0 < x < 100.0
    -2^31 <= n <= 2^31-1
    -10^4 <= xn <= 10^4
*/
func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1.0
	case n < 0:
		x = 1.0 / x
		n *= -1
	case n > 0:
	}
	return myPow2Iter(x, n)
}

func myPow2(x float64, n int) float64 {
	switch {
	case n == 1:
		return x
	case n%2 == 1:
		pn2 := myPow2(x, n>>1)
		return x * pn2 * pn2
	case n%2 == 0:
		pn2 := myPow2(x, n>>1)
		return pn2 * pn2
	default:
		panic("impossible case")
	}
}
func myPow2Iter(x float64, n int) float64 {
	ret := x
	i := 2
	for ; i <= n; i <<= 1 {
		ret *= ret
	}
	if i > n {
		i >>= 1
		for ; i < n; i++ {
			ret *= x
		}
	}
	return ret
}
