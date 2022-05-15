package conclusion

/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.

Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.

Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.

Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.

Constraints:

    1 <= nums.length <= 104
    -231 <= nums[i] <= 231 - 1

Follow up: Can you find an O(n) solution?
*/
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		panic("invalid input")
	}
	max := make([]int, 1, 3)
	max[0] = nums[0]
L:
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		//check low
		if v <= max[0] && len(max) == 3 {
			continue
		}
		//check presense in max
		//if v <= max[len(max)-1] && v >= max[0] {
		//	for _, m := range max {
		//		if m == v {
		//			continue L
		//		}
		//	}
		//}
		//max[0]<v<max[last]
		//insert v into proper position
		var j int
		for j = 0; j < len(max) && max[j] < v; j++ { //linear search because binary search is too slow on 3 elements
		}
		if j < len(max) && max[j] == v {
			continue L //already exist
		}
		//j is pos for insert
		if len(max) < 3 {
			max = append(max, 0)
			copy(max[j+1:], max[j:])
			max[j] = v
		} else {
			copy(max, max[1:j])
			max[j-1] = v
		}
	}
	if len(max) < 3 {
		return max[len(max)-1]
	}
	return max[0]
}
