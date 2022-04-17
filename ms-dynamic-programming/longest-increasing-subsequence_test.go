package ms_dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLIS(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			out:  4,
		},
		{
			nums: []int{0, 1, 0, 3, 2, 3},
			out:  4,
		},
		{
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			out:  1,
		},
	}

	for i, test := range tests {
		//res := lengthOfLIS_DP(test.nums)
		//res := lengthOfLIS_Arr(test.nums)
		res := lengthOfLIS_ArrBin(test.nums)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
