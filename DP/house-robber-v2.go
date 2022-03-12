package DP

func HouseRobberV2(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	case 2:
		return max(houses[0], houses[1])
	}
	sums := make([]int, len(houses))
	sums[0] = houses[0]
	sums[1] = max(houses[0], houses[1])
	for i := 2; i < len(houses); i++ {
		t1 := sums[i-1]
		t2 := sums[i-2]
		cur := houses[i]
		sums[i] = max(t1, t2+cur)
	}
	return max(sums[len(sums)-1], sums[len(sums)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
