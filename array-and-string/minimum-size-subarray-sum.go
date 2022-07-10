package array_and_string

import "math"

/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr]
of which the sum is greater than or equal to target.
If there is no such subarray, return 0 instead.
*/

func minSubArrayLenHint(target int, nums []int) int {
	var sum, left int
	ln := math.MaxInt
	for i, v := range nums {
		sum += v
		for sum >= target {
			if tmpLn := i + 1 - left; ln > tmpLn {
				ln = tmpLn
			}
			sum -= nums[left]
			left++
		}
	}
	if ln == math.MaxInt {
		return 0
	}
	return ln
}

func minSubArrayLen(target int, nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1: //вырожденный случай, нет смысла греть процессор сложной логикой
		if nums[0] >= target {
			return 1
		}
		return 0
	}

	var sum, ln int
	var lIdx, rIdx int
	lastIdx := len(nums) - 1

	//init cycle - go right until sum >= target
	for ; sum < target && rIdx <= lastIdx; rIdx++ {
		sum += nums[rIdx]
	}
	if sum >= target {
		ln = rIdx
	} else {
		return 0 //мы сделали всё что могли - вся последовательность меньше цели
	}
	rIdx--

	for {
		//calc min
		if tmpLn := rIdx - lIdx + 1; sum >= target && ln > tmpLn {
			ln = tmpLn
		}
		//make decision
		switch {
		case rIdx == lastIdx && (lIdx == rIdx || sum <= target):
			// дошли до конца последовательности
			//сложили всё что можно, но суммы не хватает или уже собрана, нт смысла дальше работать
			return ln
		case sum > target && lIdx < rIdx: //try to optimize
			sum -= nums[lIdx]
			lIdx++
		case sum < target && rIdx < lastIdx, lIdx == rIdx: //need more points
			rIdx++
			sum += nums[rIdx]
		case sum == target && rIdx < lastIdx: //shift window
			sum -= nums[lIdx]
			lIdx++
			rIdx++
			sum += nums[rIdx]
		default:
			panic("impossible case")
		}
	}

}
