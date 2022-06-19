package array_and_string

/*
Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal
to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left.
This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.
*/

func pivotIndexHint(nums []int) int {
	var sum, lSum int
	for _, v := range nums {
		sum += v
	}
	for i, v := range nums {
		rSum := sum - lSum - v
		if lSum == rSum {
			return i
		}
		lSum += v
	}
	return -1
}
func pivotIndex(nums []int) int {
	lIdx, rIdx := 0, len(nums)-1
	lSum, rSum := 0, 0
	for rIdx != lIdx {
		if lSum+nums[lIdx] > rSum+nums[rIdx] {
			rSum += nums[rIdx]
			rIdx--
		} else {
			lSum += nums[lIdx]
			lIdx++
		}
	}
	if lSum != rSum {
		return -1
	}
	return lIdx
}
