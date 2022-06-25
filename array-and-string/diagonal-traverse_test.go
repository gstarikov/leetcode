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
		{ //0
			in:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			out: []int{1, 2, 3, 5, 4, 6, 7, 8},
		},
		{ //1
			in:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			out: []int{1, 2, 5, 6, 3, 4, 7, 8},
		},
		{ //2
			in:  [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			out: []int{1, 2, 6, 7, 3, 4, 8, 9, 5, 10},
		},
		{ //3
			in:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			out: []int{1, 2, 3, 5, 4, 6, 7, 9, 8, 10},
		},
		{ //4
			in:  [][]int{{1, 2}, {3, 4}},
			out: []int{1, 2, 3, 4},
		},
		{ //5
			in:  [][]int{{2, 5}, {8, 4}, {0, -1}},
			out: []int{2, 5, 8, 0, 4, -1},
		},
		{ //6
			in:  [][]int{{2}, {3}},
			out: []int{2, 3},
		},
		{ //7
			in:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			out: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{ //8
			in:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			out: []int{1, 2, 5, 9, 6, 3, 4, 7, 10, 13, 14, 11, 8, 12, 15, 16},
		},
		{ //9
			in:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			out: []int{1, 2, 5, 9, 6, 3, 4, 7, 10, 11, 8, 12},
		},
		{ //10
			in:  [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}},
			out: []int{1, 2, 6, 11, 7, 3, 4, 8, 12, 16, 17, 13, 9, 5, 10, 14, 18, 19, 15, 20},
		},
	}

	for i, test := range tests {
		out := findDiagonalOrder(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}