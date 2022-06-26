package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		in  int
		out [][]int
	}{
		{
			in: 4,
			out: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
			},
		},
	}

	for i, test := range tests {
		out := generate(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
