package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in, out []int
		k       int
	}{
		{
			in:  []int{1, 1, 2},
			out: []int{1, 2},
			k:   2,
		},
		{
			in:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			out: []int{0, 1, 2, 3, 4},
			k:   2,
		},
	}

	for i, test := range tests {
		k := removeDuplicates(test.in)
		assert.Equalf(t, test.out, test.in[:k], "case(%d)", i)
	}
}
