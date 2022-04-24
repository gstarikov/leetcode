package ms_other

/*
Given a string columnTitle that represents the column title as appear in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:

Input: columnTitle = "A"
Output: 1

Example 2:

Input: columnTitle = "AB"
Output: 28

Example 3:

Input: columnTitle = "ZY"
Output: 701

Constraints:

    1 <= columnTitle.length <= 7
    columnTitle consists only of uppercase English letters.
    columnTitle is in the range ["A", "FXSHRXW"].


*/

func getC2N(c byte) int {
	return int(c) - 'A' + 1
}

func titleToNumber(columnTitle string) (acc int) {
	for _, v := range columnTitle {
		acc *= 26
		c := getC2N(byte(v))
		acc += c
	}
	return acc
}
