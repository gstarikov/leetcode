package recursion_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairsNoRecursion(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			in:  []int{1, 2, 3, 4},
			out: []int{2, 1, 4, 3},
		},
	}

	for i, test := range tests {
		head := SliceToList(test.in)
		//head = swapPairsNoRecursion
		head = swapPairsRecursion(head)
		out := ListToSlice(head)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}

func ListToSlice(lst *ListNode) []int {
	res := []int{}
	for ; lst != nil; lst = lst.Next {
		res = append(res, lst.Val)
	}
	return res
}

func SliceToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	res := &ListNode{
		Val: arr[0],
	}
	ptr := res
	for i := 1; i < len(arr); i++ {
		ptr.Next = &ListNode{
			Val: arr[i],
		}
		ptr = ptr.Next
	}
	return res
}
