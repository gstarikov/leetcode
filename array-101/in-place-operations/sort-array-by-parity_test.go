package in_place_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			in:  []int{3, 1, 2, 4},
			out: []int{4, 2, 1, 3},
		},
	}

	for i, test := range tests {
		out := sortArrayByParity(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
