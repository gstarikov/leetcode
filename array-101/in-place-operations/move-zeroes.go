package in_place_operations

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]

Constraints:

    1 <= nums.length <= 104
    -231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/
func moveZeroes(nums []int) {
	var nonZeroIdx, idx int
	//skip first non zeroes
	for ; idx < len(nums) && nums[idx] != 0; idx++ {
	}
	nonZeroIdx = idx
	//move non zeroes
	for idx < len(nums) {
		//skip zero seq
		for ; idx < len(nums) && nums[idx] == 0; idx++ {
		}
		//find last nonZero
		for ; idx < len(nums) && nums[idx] != 0; idx++ {
			nums[nonZeroIdx] = nums[idx]
			//nums[idx] = 0 //it wll slow down about +20-25%
			nonZeroIdx++
		}
	}
	//fill up tail with zeroes. it will speed up upto -20%
	for i := nonZeroIdx; i < len(nums); i++ {
		nums[i] = 0
	}
}
