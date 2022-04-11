package ms_sorting_and_searching

/*
Given an array nums with n objects colored red, white, or blue,
sort them in-place so that objects of the same color are adjacent,
with the colors in the order red, white, and blue.
We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
You must solve this problem without using the library's sort function.

Constraints:
    n == nums.length
    1 <= n <= 300
    nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/
func sortColors(nums []int) {
	var buf [3]int
	for _, v := range nums {
		buf[v]++
	}

	var pos int
	for color, cnt := range buf {
		for ; cnt > 0; cnt-- {
			nums[pos] = color
			pos++
		}
	}
}
func sortColors2(nums []int) {
	p0, cur, p2 := 0, 0, len(nums)-1
	for cur <= p2 {
		switch nums[cur] {
		case 0:
			swap(&nums[cur], &nums[p0])
			cur++
			p0++
		case 2:
			swap(&nums[cur], &nums[p2])
			p2--
		case 1:
			cur++
		}
	}
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
