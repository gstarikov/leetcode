package DP

func DeleteAndEarnV3(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	mNums := make(map[int]int)

	// prepare map
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		val := mNums[v]
		mNums[v] = val + v
		if mn > v {
			mn = v
		}
		if mx < v {
			mx = v
		}
	}

	//do processing
	var sum [2]int // moving window
	sum[0] = mNums[mn]
	sum[1] = max(mNums[mn+1], mNums[mn])
	for i := mn + 2; i <= mx; i++ {
		cur, _ := mNums[i]
		cur = max(sum[1], sum[0]+cur)
		sum[0] = sum[1]
		sum[1] = cur
	}
	return max(sum[0], sum[1])
}
