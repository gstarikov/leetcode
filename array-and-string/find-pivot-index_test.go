package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		in    []int
		pivot int
	}{
		{
			in:    []int{1, 7, 3, 6, 5, 6},
			pivot: 3,
		},
		{
			in:    []int{1, 2, 3},
			pivot: -1,
		},
		{
			in:    []int{2, 1, -1},
			pivot: 0,
		},
		{
			in:    []int{-1, -1, -1, -1, -1, 0},
			pivot: 2,
		},
	}
	for i, test := range tests {
		p := pivotIndex(test.in)
		assert.Equalf(t, test.pivot, p, "case(%d)", i)
	}
}
