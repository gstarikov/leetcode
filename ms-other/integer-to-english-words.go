package ms_other

/*
Convert a non-negative integer num to its English words representation.

Example 1:

Input: num = 123
Output: "One Hundred Twenty Three"

Example 2:

Input: num = 12345
Output: "Twelve Thousand Three Hundred Forty Five"

Example 3:

Input: num = 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

Constraints:

    0 <= num <= 231 - 1
*/

//switch is faster that map on short lists
var n2word = []string{
	"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
}

var t2word = []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var res string
	var pow = 1
	for ; ; pow *= 1000 {
		//write value of pow
		p := num / pow
		if p == 0 {
			break
		}
		n := p % 1000
		if n == 0 {
			continue
		}
		//write pow
		switch pow {
		case 1:
		case 1000:
			res = "Thousand " + res
		case 1e6:
			res = "Million " + res
		case 1e9:
			res = "Billion " + res
		}

		m := n % 100
		l := n / 100
		if m > 0 {
			res = writeNum(m) + " " + res
		}
		if l > 0 {
			res = "Hundred " + res
			res = writeNum(l) + " " + res
		}
	}
	return res[:len(res)-1]
}

func writeNum(n int) string {
	switch {
	case n <= 19:
		return n2word[n-1]
	default: // 19< X < 100
		m := n / 10
		var res string
		if m > 0 {
			res = t2word[m-2]
		}
		n %= 10
		if n > 0 {
			res = res + " " + n2word[n-1]
		}
		return res
	}
}
