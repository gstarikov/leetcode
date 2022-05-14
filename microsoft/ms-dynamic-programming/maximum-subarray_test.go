package ms_dynamic_programming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		sum  int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			sum:  6,
		},
		{
			nums: []int{1},
			sum:  1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			sum:  23,
		},
	}

	for i, test := range tests {
		res := maxSubArray(test.nums)
		assert.Equalf(t, test.sum, res, "case(%d)", i)
	}
}
