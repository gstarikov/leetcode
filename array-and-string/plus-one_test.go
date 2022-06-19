package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits, res []int
	}{
		{
			digits: []int{1, 2, 3},
			res:    []int{1, 2, 4},
		},
		{
			digits: []int{9},
			res:    []int{1, 0},
		},
	}

	for i, test := range tests {
		r := plusOne(test.digits)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
