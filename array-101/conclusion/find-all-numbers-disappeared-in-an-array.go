package conclusion

func findDisappearedNumbers(nums []int) []int {
	//idea from hint
	//mark numbers at index
	for _, v := range nums {
		if v < 0 {
			v *= -1
		}
		idx := v - 1
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}
	//check numbers
	ret := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		num := i
		idx := i - 1
		if nums[idx] > 0 {
			ret = append(ret, num)
		}
	}
	return ret
}

func findDisappearedNumbersAddMemory(nums []int) []int {
	exist := map[int]struct{}{}
	for _, v := range nums {
		exist[v] = struct{}{}
	}
	ret := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if _, ok := exist[i]; !ok {
			ret = append(ret, i)
		}
	}
	return ret
}
