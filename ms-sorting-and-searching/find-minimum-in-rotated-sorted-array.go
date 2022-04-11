package ms_sorting_and_searching

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
For example, the array nums = [0,1,2,4,5,6,7] might become:
    [4,5,6,7,0,1,2] if it was rotated 4 times.
    [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Constraints:
    n == nums.length
    1 <= n <= 5000
    -5000 <= nums[i] <= 5000
    All the integers of nums are unique.
    nums is sorted and rotated between 1 and n times.
*/

func findMin(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[l]
	}

	for l <= r {
		mid := (l + r) >> 1
		switch {
		case nums[mid] > nums[mid+1]:
			return nums[mid+1]
		case nums[mid-1] > nums[mid]:
			return nums[mid]
		case nums[mid] > nums[0]:
			l = mid + 1
		default:
			r = mid - 1
		}
	}
	return -1
}
