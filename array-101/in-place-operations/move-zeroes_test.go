package in_place_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			in:  []int{0, 1, 0, 3, 12},
			out: []int{1, 3, 12, 0, 0},
		},
		{
			in:  []int{0},
			out: []int{0},
		},
		{
			in:  []int{1},
			out: []int{1},
		},
		{
			in:  []int{1, 0, 1},
			out: []int{1, 1, 0},
		},
	}

	for i, test := range tests {
		moveZeroes(test.in)
		assert.Equalf(t, test.out, test.in, "case(%d)", i)
	}
}
