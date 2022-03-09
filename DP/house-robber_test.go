package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		houses  []int
		bestVec []int
		bestSum int
	}{
		{
			houses:  []int{},
			bestVec: nil,
			bestSum: 0,
		},
		{
			houses:  []int{0},
			bestVec: []int{0},
			bestSum: 0,
		},
		{
			houses:  []int{1},
			bestVec: []int{0},
			bestSum: 1,
		},
		{
			houses:  []int{2, 7, 9, 3, 1},
			bestVec: []int{0, 2, 4},
			bestSum: 12,
		},
		{
			houses:  []int{2, 7, 9, 3},
			bestVec: []int{0, 2},
			bestSum: 11,
		},
		{
			houses:  []int{7, 9, 1},
			bestVec: []int{1},
			bestSum: 9,
		},
		{
			houses:  []int{7, 9, 3},
			bestVec: []int{0, 2},
			bestSum: 10,
		},
	}

	for i, v := range tests {
		r := HouseRobber(v.houses)
		assert.Equalf(t, v.bestSum, r.sum, "case(%d)", i)
		assert.Equalf(t, v.bestVec, r.sequence, "case(%d)", i)
	}
}
