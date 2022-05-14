package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2, res []int
	}{
		{
			l1:  []int{2, 4, 3},
			l2:  []int{5, 6, 4},
			res: []int{7, 0, 8},
		},
		{
			l1:  []int{0},
			l2:  []int{0},
			res: []int{0},
		},
		{
			l1:  []int{9, 9, 9, 9, 9, 9, 9},
			l2:  []int{9, 9, 9, 9},
			res: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for i, test := range tests {
		l1 := makeList(test.l1)
		l2 := makeList(test.l2)
		r := AddTwoNumbers(l1, l2)
		res := listToSlice(r)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

func listToSlice(l *ListNode) []int {
	ret := []int{}
	for ; l != nil; l = l.Next {
		ret = append(ret, l.Val)
	}
	return ret
}
