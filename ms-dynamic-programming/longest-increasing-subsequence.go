package ms_dynamic_programming

import "sort"

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:

    1 <= nums.length <= 2500
    -10^4 <= nums[i] <= 10^4

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

func lengthOfLIS_DP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				newPath := dp[j] + 1
				oldPath := dp[i]
				if newPath > oldPath {
					dp[i] = newPath
				}
			}
		}

	}

	var longest int
	for _, v := range dp {
		if v > longest {
			longest = v
		}
	}

	return longest
}

func lengthOfLIS_Arr(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sub := make([]int, 0, len(nums))
	sub = append(sub, nums[0])

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		switch {
		case num > sub[len(sub)-1]:
			sub = append(sub, num)
		default:
			for j := range sub { // replace with bin search ?
				if sub[j] >= num {
					sub[j] = num
					break
				}
			}
		}
	}
	return len(sub)
}

func lengthOfLIS_ArrBin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sub := make([]int, 0, len(nums))
	sub = append(sub, nums[0])

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		switch {
		case num > sub[len(sub)-1]:
			sub = append(sub, num)
		default:
			idx := sort.SearchInts(sub, num)
			sub[idx] = num
		}
	}

	return len(sub)
}
