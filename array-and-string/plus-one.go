package array_and_string

/*
You are given a large integer represented as an integer array digits,
where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.
*/

func plusOne(digits []int) []int {
	var over int
	over = 1 //increment value
	for i := len(digits) - 1; i >= 0 && over != 0; i-- {
		digits[i] += over
		if digits[i] >= 10 {
			over = digits[i] / 10
			digits[i] %= 10
		} else {
			over = 0
			break
		}
	}
	if over > 0 {
		digits = append([]int{over}, digits...)
	}
	return digits
}
