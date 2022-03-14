package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrapRainWater(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{4, 2, 0, 3, 2, 5},
			out: 9,
		},
		{
			in:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			out: 6,
		},
	}
	for i, test := range tests {
		r := TrapRainWater2(test.in)
		assert.Equalf(t, test.out, r, "case(%d)", i)
	}
}
