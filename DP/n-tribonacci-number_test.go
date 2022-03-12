package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTribonacci(t *testing.T) {
	tests := []struct {
		n, result int
	}{
		{ //0
			n:      0,
			result: 0,
		},
		{ //1
			n:      1,
			result: 1,
		},
		{ //2
			n:      2,
			result: 1,
		},
		{ //3
			n:      3,
			result: 2,
		},
		{ //4
			n:      4,
			result: 4,
		},
		{ //5
			n:      5,
			result: 4 + 2 + 1,
		},
		{ //6
			n:      25,
			result: 1389537,
		},
	}
	for i, test := range tests {
		res := Tribonacci(test.n)
		assert.Equalf(t, test.result, res, "case(%d)", i)
	}
}
