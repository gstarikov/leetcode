package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers2(t *testing.T) {
	tests := []struct {
		l1, l2, res []int
	}{
		{
			l1:  []int{7, 2, 4, 3},
			l2:  []int{5, 6, 4},
			res: []int{7, 8, 0, 7},
		},
		{
			l1:  []int{2, 4, 3},
			l2:  []int{5, 6, 4},
			res: []int{8, 0, 7},
		},
		{
			l1:  []int{0},
			l2:  []int{0},
			res: []int{0},
		},
	}

	for i, test := range tests {
		l1, l2 := makeList(test.l1), makeList(test.l2)
		res := AddTwoNumbers2(l1, l2)
		resArr := listToSlice(res)
		assert.Equalf(t, test.res, resArr, "case(%d)", i)
	}
}
