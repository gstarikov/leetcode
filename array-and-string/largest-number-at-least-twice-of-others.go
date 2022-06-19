package array_and_string

func dominantIndex(nums []int) int {
	idx, mx := 0, nums[0]
	//find the max
	for i := 1; i < len(nums); i++ {
		if mx < nums[i] {
			mx = nums[i]
			idx = i
		}
	}
	//check max
	for _, v := range nums {
		if v != mx && mx < v*2 {
			return -1
		}
	}
	return idx
}
