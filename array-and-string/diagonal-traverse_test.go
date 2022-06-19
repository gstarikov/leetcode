package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{
			in:  [][]int{{1, 2}, {3, 4}},
			out: []int{1, 2, 3, 4},
		},
		{
			in:  [][]int{{2, 5}, {8, 4}, {0, -1}},
			out: []int{2, 5, 8, 0, 4, -1},
		},
		{
			in:  [][]int{{2}, {3}},
			out: []int{2, 3},
		},
		{
			in:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			out: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			in:  [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}},
			out: []int{1, 2, 6, 11, 7, 3, 4, 8, 12, 16, 17, 13, 9, 5, 10, 14, 18, 19, 15, 20},
		},
	}

	for i, test := range tests {
		out := findDiagonalOrder(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
