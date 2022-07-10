package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums, res []int
		k         int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			res:  []int{5, 6, 7, 1, 2, 3, 4},
			k:    3,
		}, {
			nums: []int{-1, -100, 3, 99},
			res:  []int{3, 99, -1, -100},
			k:    2,
		},
	}

	f := []func([]int, int){
		rotateTrivial,
		rotateTrivial2,
		rotateTrivial3,
		rotate,
		rotate2,
	}

	for i, test := range tests {
		for _, fn := range f {
			src := make([]int, len(test.nums))
			copy(src, test.nums)
			fn(src, test.k)
			assert.Equalf(t, test.res, src, "case(%d) function(%s)",
				i, runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		}
	}
}

func BenchmarkRotateArrayTrivial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//rotateTrivial([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		//rotateTrivial2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		//rotateTrivial3([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		//rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
		rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}
