package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    int
	}{
		{ //0
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			res:    4,
		},
		{ //1
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			res:    -1,
		},
		{ //2
			nums:   []int{1},
			target: 0,
			res:    -1,
		},
		{ //3
			nums:   []int{1, 3},
			target: 0,
			res:    -1,
		},
	}

	for i, test := range tests {
		res := search(test.nums, test.target)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}

}
