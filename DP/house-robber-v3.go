package DP

func HouseRobberV3(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	case 2:
		return max(houses[0], houses[1])
	}
	// sums := make([]int, len(houses)) // its like moving window
	t2 := houses[0] // t2 = t minus 2 = sum[i-2]
	t1 := houses[1]
	t1 = max(t2, t1)
	for i := 2; i < len(houses); i++ {
		cur := houses[i]
		cur = max(t1, t2+cur)
		t2, t1 = t1, cur // move window
	}
	return max(t1, t2)
}
