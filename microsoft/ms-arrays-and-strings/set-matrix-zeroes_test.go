package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		in, out [][]int
	}{
		{in: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, out: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{in: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, out: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{in: [][]int{{1, 0}}, out: [][]int{{0, 0}}},
		{in: [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}, out: [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
	}

	for i, test := range tests {
		SetZeroes(test.in)
		assert.Equalf(t, test.out, test.in, "case(%d)", i)
	}
}
