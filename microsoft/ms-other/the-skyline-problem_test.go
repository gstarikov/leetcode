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
		{ //0
			buildings: [][]int{{15, 20, 10}, {19, 24, 8}},
			res:       [][]int{{15, 10}, {20, 8}, {24, 0}},
		},
		{ //1
			buildings: [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			res:       [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{ //2
			buildings: [][]int{{0, 2, 3}, {2, 5, 3}},
			res:       [][]int{{0, 3}, {5, 0}},
		},
		{ //3
			buildings: [][]int{{0, 2147483647, 2147483647}},
			res:       [][]int{{0, 2147483647}, {2147483647, 0}},
		},
		{ //4
			buildings: [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}},
			res:       [][]int{{1, 3}, {2, 0}},
		},
		{ //5
			buildings: [][]int{{0, 3, 3}, {1, 5, 3}, {2, 4, 3}, {3, 7, 3}},
			res:       [][]int{{0, 3}, {7, 0}},
		},
		{ //6
			buildings: [][]int{
				{1, 38, 219},
				{2, 19, 228}, {2, 64, 106},
				{3, 80, 65}, {3, 84, 8},
				{4, 12, 8}, {4, 25, 14}, {4, 46, 225}, {4, 67, 187},
				{5, 36, 118}, {5, 48, 211}, {5, 55, 97},
				{6, 42, 92}, {6, 56, 188},
				{7, 37, 42}, {7, 49, 78}, {7, 84, 163},
				{8, 44, 212},
				{9, 42, 125}, {9, 85, 200}, {9, 100, 74},
				{10, 13, 58},
				{11, 30, 179},
				{12, 32, 215}, {12, 33, 161}, {12, 61, 198},
				{13, 38, 48}, {13, 65, 222},
				{14, 22, 1},
				{15, 70, 222},
				{16, 19, 196}, {16, 24, 142}, {16, 25, 176}, {16, 57, 114},
				{18, 45, 1},
				{19, 79, 149},
				{20, 33, 53}, {21, 29, 41}, {23, 77, 43}, {24, 41, 75}, {24, 94, 20}, {27, 63, 2}, {31, 69, 58}, {31, 88, 123}, {31, 88, 146}, {33, 61, 27}, {35, 62, 190}, {35, 81, 116}, {37, 97, 81}, {38, 78, 99}, {39, 51, 125}, {39, 98, 144}, {40, 95, 4}, {45, 89, 229}, {47, 49, 10}, {47, 99, 152}, {48, 67, 69}, {48, 72, 1}, {49, 73, 204}, {49, 77, 117}, {50, 61, 174}, {50, 76, 147}, {52, 64, 4}, {52, 89, 84}, {54, 70, 201}, {57, 76, 47}, {58, 61, 215}, {58, 98, 57}, {61, 95, 190}, {66, 71, 34}, {66, 99, 53}, {67, 74, 9}, {68, 97, 175}, {70, 88, 131}, {74, 77, 155}, {74, 99, 145}, {76, 88, 26}, {82, 87, 40}, {83, 84, 132}, {88, 99, 99}},
			res: [][]int{{1, 219}, {2, 228}, {19, 225}, {45, 229}, {89, 190}, {95, 175}, {97, 152}, {99, 74}, {100, 0}},
		},
	}

	for i, test := range tests {
		res := getSkylineTry3(test.buildings)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

func TestHelperMerge(t *testing.T) {
	tests := []struct {
		in, out              []point
		lPos, l, r, h, prevH int
	}{
		{ //0
			in:    []point{{1, 1}, {2, 0}},
			out:   []point{{0, 2}, {2, 0}},
			lPos:  0,
			l:     0,
			r:     2,
			h:     2,
			prevH: 0,
		},
		{ //1
			in:    []point{{1, 2}, {2, 3}, {3, 1}, {4, 2}, {5, 0}},
			out:   []point{{1, 2}, {2, 3}, {3, 2}, {5, 0}},
			lPos:  0,
			l:     1,
			r:     4,
			h:     2,
			prevH: 2,
		},
	}

	for i, test := range tests {
		res := helperMerge(test.lPos, test.l, test.r, test.h, test.prevH, test.in)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}

func TestPQueue(t *testing.T) {
	tests := []struct {
		remove   bool
		pnt      point
		internal []point
	}{
		{ //0
			remove:   true,
			internal: nil,
		},
		{ //1
			remove:   false,
			pnt:      point{1, 1},
			internal: []point{{1, 1}},
		},
		{ //2
			remove:   true,
			internal: []point{},
		},
		{ //3
			remove:   false,
			pnt:      point{1, 1},
			internal: []point{{1, 1}},
		},
		{ //4
			remove:   false,
			pnt:      point{1, 1},
			internal: []point{{1, 1}},
		},
		{ //5
			remove:   false,
			pnt:      point{1, 2},
			internal: []point{{1, 2}},
		},
		{ //6
			remove:   false,
			pnt:      point{2, 1},
			internal: []point{{1, 2}, {2, 1}},
		},
		{ //7
			remove:   false,
			pnt:      point{3, 2},
			internal: []point{{3, 2}},
		},
		{ //8
			remove:   false,
			pnt:      point{5, 3},
			internal: []point{{5, 3}},
		},
		{ //9
			remove:   false,
			pnt:      point{3, 3},
			internal: []point{{5, 3}},
		},
		{ //10
			remove:   false,
			pnt:      point{3, 2},
			internal: []point{{5, 3}},
		},
		{ //11
			remove:   false,
			pnt:      point{10, 2},
			internal: []point{{5, 3}, {10, 2}},
		},
		{ //12
			remove:   true,
			internal: []point{{10, 2}},
		},
		{ //13
			remove:   false,
			pnt:      point{20, 2},
			internal: []point{{20, 2}},
		},
	}

	pq := pQueue{}
	for i, test := range tests {
		if test.remove {
			pq.Remove()
		} else {
			pq.Add(test.pnt)
		}
		assert.Equalf(t, test.internal, pq.queu, "case(%d)", i)
	}
}
