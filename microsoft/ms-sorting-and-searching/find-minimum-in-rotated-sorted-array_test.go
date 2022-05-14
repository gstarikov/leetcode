package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			out:  1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			out:  0,
		},
		{
			nums: []int{11, 13, 15, 17},
			out:  11,
		},
	}

	for i, test := range tests {
		res := findMin(test.nums)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
