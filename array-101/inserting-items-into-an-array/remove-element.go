package inserting_items_into_an_array

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}
