package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			in:  []int{2, 0, 2, 1, 1, 0},
			out: []int{0, 0, 1, 1, 2, 2},
		},
		{
			in:  []int{2, 0, 1},
			out: []int{0, 1, 2},
		},
		{
			in:  []int{1, 2, 0},
			out: []int{0, 1, 2},
		},
	}

	for i, test := range tests {
		sortColors2(test.in)
		assert.Equalf(t, test.out, test.in, "case(%d)", i)
	}
}
