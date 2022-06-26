package array_and_string

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	if len(nums) < 2 {
		return 0 //must be panic
	}
	sort.Ints(nums)
	sum := mMin(nums[0], nums[1])
	for i := 2; i < len(nums); i += 2 {
		sum += mMin(nums[i], nums[i+1])
	}
	return sum
}

func mMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
