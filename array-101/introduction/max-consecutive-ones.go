package introduction

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/

func findMaxConsecutiveOnes(nums []int) int {
	var maxN int
	var i, n int
	n = len(nums)
	for i < n {
		//skip zeroes
		for ; i < n && nums[i] == 0; i++ {
		}
		//calc ones
		var curN = i
		for ; i < n && nums[i] == 1; i++ {
			//curN++
		}
		curN = i - curN
		if curN > maxN {
			maxN = curN
		}
	}
	return maxN
}
