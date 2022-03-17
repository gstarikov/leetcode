package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists2(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{
			in: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			out: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			in:  [][]int{},
			out: []int{},
		},
		{
			in:  [][]int{{}},
			out: []int{},
		},
		{
			in: [][]int{
				{1, 2, 2},
				{1, 1, 2},
			},
			out: []int{1, 1, 1, 2, 2, 2},
		},
	}

	for i, test := range tests {
		var in []*ListNode
		for _, v := range test.in {
			in = append(in, makeList(v))
		}
		res := MergeKLists(in)
		resArr := listToSlice(res)
		assert.Equalf(t, test.out, resArr, "case(%d)", i)
	}
}
