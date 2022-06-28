package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		out     []int
	}{
		{ //0
			numbers: []int{2, 7, 11, 15},
			target:  9,
			out:     []int{1, 2},
		},
		{ //1
			numbers: []int{2, 3, 4},
			target:  6,
			out:     []int{1, 3},
		},
		{ //2
			numbers: []int{-1, 0},
			target:  -1,
			out:     []int{1, 2},
		},
		{ //3
			numbers: []int{0, 0, 3, 4},
			target:  0,
			out:     []int{1, 2},
		},
	}

	for i, test := range tests {
		out := twoSum(test.numbers, test.target)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
