package array_and_string

import "sort"

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.
*/

func twoSum(numbers []int, target int) []int {
	/*
		1 - находим точку равную target-первое число
		2 - идём через 2 индекса с первого числа и с найденного и ищем сумму
	*/

	rVal := target - numbers[0]
	rIdx := sort.SearchInts(numbers, rVal)
	if rIdx == len(numbers) {
		rIdx--
	}

	//shift rVal till all eq vals
	if rIdx+1 < len(numbers) && numbers[rIdx] == numbers[rIdx+1] {
		rIdx++
	}

	lVal := target - numbers[rIdx]
	lIdx := sort.SearchInts(numbers[:rIdx], lVal)
	if lIdx >= rIdx {
		panic("impossible case")
	}
L:
	for lIdx < rIdx {
		sum := numbers[rIdx] + numbers[lIdx]
		switch {
		case sum == target:
			break L
		case sum < target:
			lIdx++
		case sum > target:
			rIdx--
		}
	}
	return []int{lIdx + 1, rIdx + 1}
}
