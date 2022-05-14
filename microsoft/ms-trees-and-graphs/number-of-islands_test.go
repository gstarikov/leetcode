package ms_trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		in  [][]byte
		res int
	}{
		{in: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}},
			res: 1},
		{in: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}},
			res: 3},
	}
	for i, test := range tests {
		res := NumIslands(test.in)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
