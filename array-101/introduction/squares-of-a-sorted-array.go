package introduction

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	//analyze zero point
	l, r := nums[0], nums[len(nums)-1]
	ln := len(nums)
	switch {
	case l >= 0 && r >= 0: //dont rotate array
		for i := range nums {
			nums[i] *= nums[i]
		}
	case l < 0 && r > 0: //there is rotate point
		//must allocate array
		ret := make([]int, len(nums))
		pos := len(ret) - 1

		for l, r = 0, len(nums)-1; l <= r; pos-- {
			ll, rr := abs(int64(nums[l])), abs(int64(nums[r]))
			//sqL, sqR := nums[l]*nums[l], nums[r]*nums[r]
			if ll <= rr {
				ret[pos] = int(rr * rr)
				r--
			} else {
				ret[pos] = int(ll * ll)
				l++
			}
		}
		nums = ret
	case l < 0 && r <= 0: //do reverse array
		for l, r = 0, ln-1; l < r; l, r = l+1, r-1 {
			tmp := nums[r]
			nums[r] = nums[l]
			nums[l] = tmp
			nums[l] *= nums[l]
			nums[r] *= nums[r]
		}
		if l == r {
			nums[l] *= nums[l]
		}
	default:
		panic("impossible case")
	}
	return nums
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
