package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1, l2, out []int
	}{
		{
			l1:  []int{1, 2, 4},
			l2:  []int{1, 3, 4},
			out: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:  []int{},
			l2:  []int{},
			out: []int{},
		},
		{
			l1:  []int{},
			l2:  []int{0},
			out: []int{0},
		},
	}

	for i, test := range tests {
		l1, l2 := makeList(test.l1), makeList(test.l2)
		res := MergeTwoLists(l1, l2)
		resArr := listToSlice(res)
		assert.Equalf(t, test.out, resArr, "case(%d)", i)
	}
}
