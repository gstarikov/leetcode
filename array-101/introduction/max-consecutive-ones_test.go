package introduction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			res:  3,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1},
			res:  2,
		},
	}

	for i, test := range tests {
		res := findMaxConsecutiveOnes(test.nums)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
