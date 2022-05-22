package recursion_i

/*
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row.
Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01,
and each occurrence of 1 with 10.

    For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110.

Given two integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

Constraints:
    1 <= n <= 30
    1 <= k <= 2^(n - 1)
*/

func kthGrammar(n int, k int) int {
	var cur, prev []byte
	cur = []byte{0}
	for i := 2; i <= n; i++ {
		tmp := prev[:0]
		prev = cur
		cur = tmp

		for _, v := range prev {
			switch v {
			case 0:
				cur = append(cur, 0, 1)
			case 1:
				cur = append(cur, 1, 0)
			default:
				panic("impossible case")
			}
		}
	}
	return int(cur[k-1])
}

/*
на каждой итерации:
после 0 подписываем 1
после 1 подписываем 0

1 -> 0
2 -> 01
3 -> 0110
4 -> 01101001
5 -> 0110100110010110

f(5,7) (0) -> f(4,3) (1) (idx +1) -> f(3, 1) (0) (idx+1)

1 элемент всегда 0
2 элемент появляется из 1 +1
3 элемент появляется из 2 +0
4 элемент из 2 +1
5 элемент из 3 +0 (5+1/2
6 элемент из 3 +1
*/

func kthGrammarRec(n, k int) int {
	return kthGrammarRecInt(k)
}

func kthGrammarRecInt(k int) int {
	//base case
	if k == 1 {
		return 0
	}
	prev := kthGrammarRecInt((k + 1) >> 1)
	k2 := k % 2
	switch { //use switch instead of slice to reduce gc
	case prev == 0 && k2 == 0:
		return 1
	case prev == 0 && k2 == 1:
		return 0
	case prev == 1 && k2 == 0:
		return 0
	case prev == 1 && k2 == 1:
		return 1
	default:
		panic("impossible case")
	}
}
