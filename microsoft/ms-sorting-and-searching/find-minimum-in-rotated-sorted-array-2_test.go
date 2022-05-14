package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMin2(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{
			nums: []int{1, 3, 5},
			out:  1,
		},
		{
			nums: []int{2, 2, 2, 0, 1},
			out:  0,
		}, {
			nums: []int{1, 1},
			out:  1,
		},
	}

	for i, test := range tests {
		res := findMin2(test.nums)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
