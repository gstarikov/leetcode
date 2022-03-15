package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{in: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{in: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, out: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for i, test := range tests {
		r := SpiralOrder(test.in)
		assert.Equalf(t, test.out, r, "case(%d)", i)
	}
}
