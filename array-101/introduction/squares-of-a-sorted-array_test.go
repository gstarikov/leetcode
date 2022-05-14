package introduction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{ //0
			in:  []int{-4, -1, 0, 3, 10},
			out: []int{0, 1, 9, 16, 100},
		},
		{ //1
			in:  []int{-7, -3, 2, 3, 11},
			out: []int{4, 9, 9, 49, 121},
		},
		{ //2
			in:  []int{-5, -3, -2, -1},
			out: []int{1, 4, 9, 25},
		},
		{ //3
			in:  []int{-5, -3, -2, 2},
			out: []int{4, 4, 9, 25},
		},
		{ //4
			in:  []int{-4, -4, -3},
			out: []int{9, 16, 16},
		},
		{ //5
			in:  []int{-1},
			out: []int{1},
		},
	}
	for i, test := range tests {
		out := sortedSquares(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}

/*
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkAbs-8   	1000000000	         1.012 ns/op  //math.Abs(float64(j))
BenchmarkAbs-8   	1000000000	         1.050 ns/op // j*j
BenchmarkAbs-8   	1000000000	         0.9734 ns/op // abs
*/
func BenchmarkAbs(b *testing.B) {
	//j := -1
	for i := 0; i < b.N; i++ {
		_ = abs(int64(i))
	}
}
