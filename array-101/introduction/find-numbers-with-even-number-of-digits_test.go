package introduction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		in  []int
		res int
	}{
		{
			in:  []int{12, 345, 2, 6, 7896},
			res: 2,
		},
		{
			in:  []int{555, 901, 482, 1771},
			res: 1,
		},
	}
	for i, test := range tests {
		res := findNumbers(test.in)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
