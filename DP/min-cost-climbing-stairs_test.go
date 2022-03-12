package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		sum  int
	}{
		{ //0
			cost: nil,
			sum:  0,
		},
		{ //1
			cost: []int{},
			sum:  0,
		},
		{ //2
			cost: []int{1},
			sum:  1,
		},
		{ //3
			cost: []int{1, 2},
			sum:  1,
		},
		{ //4
			cost: []int{1, 2, 1, 2},
			sum:  2,
		},
		{ //5
			cost: []int{1, 2, 2, 2, 1, 1, 2},
			sum:  5,
		},
		{ //6
			cost: []int{10, 15, 20},
			sum:  15,
		},
		{ //7
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			sum:  6,
		},
	}
	for i, test := range tests {
		res := MinCostClimbingStairs(test.cost)
		assert.Equalf(t, test.sum, res, "case(%d)", i)
	}
}

func BenchmarkMinCostClimbingStairs(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MinCostClimbingStairs(cost)
	}
}
