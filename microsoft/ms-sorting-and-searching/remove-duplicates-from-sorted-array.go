package ms_sorting_and_searching

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
*/

func removeDuplicates(nums []int) int {
	last := len(nums)
	var uniqCnt int
	for i := 0; i < last; i++ {
		//check duplicates
		v := nums[i]
		for j := i + 1; j < last && v == nums[j]; j++ {
			i++
		}

		nums[uniqCnt] = nums[i] // copy last of dup sequence
		uniqCnt++
	}
	return uniqCnt
}
