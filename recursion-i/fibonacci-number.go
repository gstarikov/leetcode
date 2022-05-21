package recursion_i

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	i1 := 0
	i2 := 1
	for i := 2; i < n; i++ {
		l := i1 + i2
		i1 = i2
		i2 = l
	}
	return i1 + i2
}
func fibRec(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	cache := make([]int, n)
	cache[0], cache[1] = 0, 1
	return fibRecc(cache, n-2) + fibRecc(cache, n-1)
}
func fibRecc(cache []int, n int) int {
	if n < 2 {
		return cache[n]
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n-1] = fibRecc(cache, n-1)
	cache[n-2] = fibRecc(cache, n-2)
	return cache[n-1] + cache[n-2]
}
