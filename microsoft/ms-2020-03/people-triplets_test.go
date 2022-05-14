package ms_2020_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeopleTriplets(t *testing.T) {
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
			r: 3, //was error in task definition. was 4 must be 3
		},
	}

	for i, test := range tests {
		res := CountOfTriplets(test.p, test.t)
		assert.Equalf(t, test.r, res, "case(%d) test(%+v) res(%d)", i, test, res)
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		in, res int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}

	for i, test := range tests {
		res := factorial(test.in)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

func TestCalcTripletsCombinations(t *testing.T) {
	tests := []struct {
		in, res int
	}{
		{3, 1},
		{4, 4},
		{5, 10},
	}

	for i, test := range tests {
		res := calcTripletsCombinations(3, test.in)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
