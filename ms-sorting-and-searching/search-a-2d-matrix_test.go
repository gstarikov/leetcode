package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		in     [][]int
		target int
		res    bool
	}{
		{ //0
			in: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			res:    true,
		},
		{ //1
			in: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			res:    false,
		},
		{ //2
			in: [][]int{
				{1, 3},
			},
			target: 3,
			res:    true,
		},
	}

	for i, test := range tests {
		res := searchMatrixHint(test.in, test.target)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
