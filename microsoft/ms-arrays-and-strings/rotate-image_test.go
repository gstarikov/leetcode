package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateImage(t *testing.T) {
	testt := []struct {
		in, out [][]int
	}{
		{
			in:  [][]int{{1, 2}, {3, 4}},
			out: [][]int{{3, 1}, {4, 2}},
		},
		{
			in:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			out: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			in:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			out: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for i, test := range testt {
		RotateImage(test.in)
		assert.Equalf(t, test.out, test.in, "case(%d)", i)
	}
}
