package inserting_items_into_an_array

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

    arr.length >= 3
    There exists some i with 0 < i < arr.length - 1 such that:
        arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
        arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Example 1:

Input: arr = [2,1]
Output: false

Example 2:

Input: arr = [3,5,5]
Output: false

Example 3:

Input: arr = [0,3,2,1]
Output: true

Constraints:
    1 <= arr.length <= 104
    0 <= arr[i] <= 104
*/

func validMountainArray(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	var i int
	var prev int
	prev = arr[0]
	i++
	for ; i < len(arr); i++ {
		v := arr[i]
		if v == prev {
			return false
		}
		if v < prev {
			break
		}
		prev = v
	}
	if i == len(arr) || i == 1 {
		return false
	}
	for ; i < len(arr); i++ {
		v := arr[i]
		if v >= prev {
			return false
		}
		prev = v
	}
	return true
}
