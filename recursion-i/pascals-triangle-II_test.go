package recursion_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct {
		r   int
		ret []int
	}{
		{
			r:   0,
			ret: []int{1},
		},
		{
			r:   1,
			ret: []int{1, 1},
		},
		{
			r:   2,
			ret: []int{1, 2, 1},
		},
		{
			r:   3,
			ret: []int{1, 3, 3, 1},
		},
		{
			r:   4,
			ret: []int{1, 4, 6, 4, 1},
		},
	}
	for i, test := range tests {
		ret := getRow(test.r)
		assert.Equalf(t, test.ret, ret, "case(%d)", i)
	}
}
