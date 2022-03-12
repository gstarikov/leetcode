package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAndEarn(t *testing.T) {
	//testHelperDeleteAndEarn(t, DeleteAndEarn)   // 5.8s (case 5)
	//testHelperDeleteAndEarn(t, DeleteAndEarnV2) // pass case 6 at light speed, but timeout on case 7
	testHelperDeleteAndEarn(t, DeleteAndEarnV3) // 0.00s
}

func testHelperDeleteAndEarn(t testing.TB, f func([]int) int) {
	tests := []struct {
		nums []int
		res  int
	}{
		{ //0
			nums: nil,
			res:  0,
		},
		{ //1
			nums: []int{},
			res:  0,
		},
		{ //2
			nums: []int{1},
			res:  1,
		},
		{ //3
			nums: []int{3, 4, 2},
			res:  6,
		},
		{ //4
			nums: []int{2, 2, 3, 3, 3, 4},
			res:  9,
		},
		{ //5
			nums: []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10},
			res:  37,
		},
		{ //6
			nums: []int{10, 8, 4, 2, 1, 3, 4, 8, 2, 9, 10, 4, 8, 5, 9, 1, 5, 1, 6, 8, 1, 1, 6, 7, 8, 9, 1, 7, 6, 8, 4, 5, 4, 1, 5, 9, 8, 6, 10, 6, 4, 3, 8, 4, 10, 8, 8, 10, 6, 4, 4, 4, 9, 6, 9, 10, 7, 1, 5, 3, 4, 4, 8, 1, 1, 2, 1, 4, 1, 1, 4, 9, 4, 7, 1, 5, 1, 10, 3, 5, 10, 3, 10, 2, 1, 10, 4, 1, 1, 4, 1, 2, 10, 9, 7, 10, 1, 2, 7, 5},
			res:  338,
		},
		{ //7
			nums: []int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93},
			res:  3451,
		},
	}
	for i, test := range tests {
		r := f(test.nums)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
