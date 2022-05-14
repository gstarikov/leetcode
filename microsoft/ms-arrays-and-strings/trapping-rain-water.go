package ms_arrays_and_strings

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

*/

func TrapRainWater(height []int) int {
	minH, maxH := 0, 0
	left := 0
	right := len(height) - 1
	sum := 0

	for {
		//find left border
		for ; left < len(height) && height[left] <= minH; left++ {
		}

		//find right border
		for ; right > 0 && height[right] <= minH; right-- {
		}

		if left >= right {
			break
		}

		//found new maxH
		if height[left] > height[right] { //yep so assigment because it is water, so max is min
			maxH = height[right]
		} else {
			maxH = height[left]
		}

		//calculate sum (integrate)
		// dot is filled in case of val < h (obsense of bar)
		for i := left + 1; i <= right-1; i++ {
			if height[i] < maxH {
				if height[i] > minH {
					sum += maxH - height[i]
				} else {
					sum += maxH - minH
				}
			}
		}

		//bring up mark
		minH = maxH
	}

	return sum
}

func TrapRainWater2(height []int) int {
	leftMax, rightMax := 0, 0
	left := 0
	right := len(height) - 1
	sum := 0

	for left < right {
		hl := height[left]
		hr := height[right]
		if hl < hr {
			if hl >= leftMax {
				leftMax = hl
			} else {
				sum += leftMax - hl
			}
			left++
		} else {
			if hr >= rightMax {
				rightMax = hr
			} else {
				sum += rightMax - hr
			}
			right--
		}
	}

	return sum
}
