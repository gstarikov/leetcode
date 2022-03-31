package ms_2020_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		p []int
		t int
		r int
	}{
		{ //0
			p: []int{1, 2, 3},
			t: 6,
			r: 1,
		},
		{ //1
			p: []int{1, 2, 3, 4, 5},
			t: 8,
			r: 2,
		},
		{ //2
			p: []int{1, 1, 1, 1},
			t: 3,
			r: 4,
		},
		{ //3
			p: []int{1, 2, 2, 2},
			t: 5,
			r: 4,
		},
	}

	for i, test := range tests {
		res := CountOfTriplets(test.p, test.t)
		assert.Equalf(t, test.r, res, "case(%d) test(%+v) res(%d)", i, test, res)
	}
}
