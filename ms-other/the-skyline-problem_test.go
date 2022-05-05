package ms_other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSkyline(t *testing.T) {
	tests := []struct {
		buildings [][]int
		res       [][]int
	}{
		{
			buildings: [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			res:       [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			buildings: [][]int{{0, 2, 3}, {2, 5, 3}},
			res:       [][]int{{0, 3}, {5, 0}},
		},
		{
			buildings: [][]int{{0, 2147483647, 2147483647}},
			res:       [][]int{{0, 2147483647}, {2147483647, 0}},
		},
	}

	for i, test := range tests {
		res := getSkylineVector(test.buildings)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

func TestHelperMerge(t *testing.T) {
	tests := []struct {
		in, out        [][]int
		l, r, h, prevH int
	}{
		{ //0
			in:    [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 2}, {5, 0}},
			out:   [][]int{{1, 2}, {2, 3}, {3, 2}, {5, 0}},
			l:     1,
			r:     4,
			h:     2,
			prevH: 2,
		},
	}

	for i, test := range tests {
		res := helperMerge(test.l, test.r, test.h, test.prevH, test.in)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
