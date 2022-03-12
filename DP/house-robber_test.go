package DP

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

func TestHouseRobber(t *testing.T) {
	testHelper(t, HouseRobber)
	testHelper(t, HouseRobberV2)
	testHelper(t, HouseRobberV3)
}

func testHelper(t testing.TB, f func([]int) int) {
	tests := []struct {
		houses  []int
		bestVec []int
		bestSum int
	}{
		{ //0
			houses:  []int{},
			bestVec: nil,
			bestSum: 0,
		},
		{ //1
			houses:  []int{0},
			bestVec: []int{0},
			bestSum: 0,
		},
		{ //2
			houses:  []int{1},
			bestVec: []int{0},
			bestSum: 1,
		},
		{ //3
			houses:  []int{2, 7, 9, 3, 1},
			bestVec: []int{0, 2, 4},
			bestSum: 12,
		},
		{ //4
			houses:  []int{2, 7, 9, 3},
			bestVec: []int{0, 2},
			bestSum: 11,
		},
		{ //5
			houses:  []int{7, 9, 1},
			bestVec: []int{1},
			bestSum: 9,
		},
		{ //6
			houses:  []int{7, 9, 3},
			bestVec: []int{0, 2},
			bestSum: 10,
		},
	}

	for i, v := range tests {
		r := f(v.houses)
		assert.Equalf(t, v.bestSum, r, "case(%d) func(%s)", i, GetFunctionName(f))
	}

}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
