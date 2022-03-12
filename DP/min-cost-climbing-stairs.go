package DP

/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
*/

func MinCostClimbingStairs(cost []int) int {
	switch len(cost) {
	case 0:
		return 0
	case 1:
		return cost[0]
	}
	t2 := cost[0]
	t1 := cost[1]
	for i := 2; i < len(cost); i++ {
		cur := cost[i]
		cur += min(t2, t1)
		t2, t1 = t1, cur
	}
	return min(t1, t2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
