package ms_sorting_and_searching

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function,
nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target,
return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1

Constraints:
    1 <= nums.length <= 5000
    -10^4 <= nums[i] <= 10^4
    All values of nums are unique.
    nums is an ascending array that is possibly rotated.
    -10^4 <= target <= 10^4
*/

func search(nums []int, target int) int {
	//border case
	switch {
	case len(nums) == 0:
		return -1
	case len(nums) == 1:
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		v := nums[mid]
		switch {
		case v == target:
			return mid
		case target == nums[l]: //it may help and isnt costly
			return l
		case target == nums[r]: //it may help and isnt costly
			return r
		case nums[l] < v: //[l:mid] is non rotated
			if target > nums[l] && target < v { //left
				r = mid - 1
			} else {
				l = mid + 1
			}
		case nums[l] > v: // [l:mid] is rotated
			if target > v && target < nums[r] { // [mid:r] search
				l = mid + 1
			} else {
				r = mid - 1
			}
		default:
			return -1
		}
	}
	return -1
}
