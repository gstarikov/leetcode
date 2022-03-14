package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		in, out []int
		target  int
	}{
		{
			in:     []int{2, 7, 11, 15},
			target: 9,
			out:    []int{0, 1},
		},
		{
			in:     []int{3, 2, 4},
			target: 6,
			out:    []int{1, 2},
		},
		{
			in:     []int{3, 3},
			target: 6,
			out:    []int{0, 1},
		},
	}

	for i, test := range tests {
		r := TwoSum(test.in, test.target)
		assert.Equalf(t, test.out, r, "case(%d)", i)
	}
}
