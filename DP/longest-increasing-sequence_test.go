package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncSeq(t *testing.T) {
	tests := []struct {
		seq []int
		len int
	}{
		{ //0
			seq: []int{},
			len: 0,
		},
		{ //1
			seq: []int{1},
			len: 1,
		},
		{ //2
			seq: []int{1, 1},
			len: 1,
		},
		{ //3
			seq: []int{1, 2},
			len: 2,
		},
		{ //4
			seq: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			len: 9,
		},
		{ //5
			seq: []int{1, 2, 2, 3, 3, 3, 4},
			len: 4,
		},
		{ //6
			seq: []int{1, 2, 1},
			len: 2,
		},
		{ //7
			seq: []int{1, 2, 1, 2},
			len: 2,
		},
		{ //8
			seq: []int{1, 2, 3, 1, 2},
			len: 3,
		},
		{ //9
			seq: []int{1, 2, 1, 2, 3},
			len: 3,
		},
		{ //10
			seq: []int{1, 2, 3, 1, 2, 2, 4, 4, 4, 5, 3, 6, 7, 3},
			len: 7,
		},
	}
	for i, test := range tests {
		res := LongestIncSeq(test.seq)
		assert.Equalf(t, test.len, res, "case(%d)", i)
	}
}
