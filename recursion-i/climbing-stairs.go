package recursion_i

/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
    1 <= n <= 45

1

11
2

111
21
12

1111
121
211
112
22
*/

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	cache := make([]int, n+1)
	return climb(cache, n)
}

func climb(cache []int, n int) int {
	if n <= 2 {
		return n
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n-1] = climb(cache, n-1)
	cache[n-2] = climb(cache, n-2)
	return cache[n-1] + cache[n-2]
}
