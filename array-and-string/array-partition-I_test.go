package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		arr []int
		res int
	}{
		{
			arr: []int{1, 4, 3, 2},
			res: 4,
		},
		{
			arr: []int{6, 2, 6, 5, 1, 2},
			res: 9,
		},
	}
	for i, test := range tests {
		res := arrayPairSum(test.arr)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
