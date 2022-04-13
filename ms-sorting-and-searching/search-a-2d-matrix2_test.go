package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix2(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		res    bool
	}{
		{ //0
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			res:    true,
		},
		{ //1
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			res:    false,
		},
		{ //2
			matrix: [][]int{
				{-5},
			},
			target: -5,
			res:    true,
		},
		{ //3
			matrix: [][]int{
				{1, 1},
			},
			target: 1,
			res:    true,
		},
		{ //4
			matrix: [][]int{
				{-1, 3},
			},
			target: 3,
			res:    true,
		},
	}
	for i, test := range tests {
		res := searchMatrix2hint(test.matrix, test.target)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
