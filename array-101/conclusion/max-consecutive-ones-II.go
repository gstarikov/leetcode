package conclusion

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.

Example 1:

Input: nums = [1,0,1,1,0]
Output: 4
Explanation: Flip the first zero will get the maximum number of consecutive 1s. After flipping, the maximum number of consecutive 1s is 4.

Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 4

Constraints:

    1 <= nums.length <= 105
    nums[i] is either 0 or 1.

Follow up: What if the input numbers come in one by one as an infinite stream?
In other words, you can't store all numbers coming from the stream as it's too large to hold in memory.
Could you solve it efficiently?
*/

func findMaxConsecutiveOnes(nums []int) int {
	//flipping= один ноль в последовательности можно заменить на 1
	// соответственно ищем последоваельность включающую не более 1 нуля
	var l, r int
	var longest int
	var numZeroes int

	for r < len(nums) {
		if nums[r] == 0 {
			numZeroes++
		}
		for numZeroes == 2 {
			if nums[l] == 0 {
				numZeroes--
			}
			l++
		}
		//new correct window
		curSeq := r - l + 1
		if curSeq > longest {
			longest = curSeq
		}
		r++
	}

	return longest
}
