package DP

/*
You are given an integer array nums. You want to maximize the number of points you get by performing the following operation any number of times:
Pick any nums[i] and delete it to earn nums[i] points. Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
Return the maximum number of points you can earn by applying the above operation some number of times.
*/

func DeleteAndEarn(nums []int) int {
	// check border conditions
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	lastMax := 0
	for _, v := range nums {
		if v == 0 {
			continue
		}
		// copy & delete
		vect, sum := copyAndDelete(nums, v)
		val := DeleteAndEarn(vect) + sum
		if val > lastMax {
			lastMax = val
		}
	}
	return lastMax
}

func copyAndDelete(nums []int, e int) (ret []int, sum int) {
	ret = make([]int, 0, len(nums)-1) //-1 because we delete at least one elem
	e1 := e - 1
	e2 := e + 1
	for _, v := range nums {
		switch v {
		case e1, e2:
			continue
		case e:
			sum += e
		default:
			ret = append(ret, v)
		}
	}
	return ret, sum
}
