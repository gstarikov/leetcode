package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		nums        []int
		target, res int
	}{
		{
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			res:    2,
		},
		{
			nums:   []int{1, 4, 4},
			target: 4,
			res:    1,
		},
		{
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			res:    0,
		},
	}

	for i, test := range tests {
		res := minSubArrayLenHint(test.target, test.nums)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

//9 ns/op
func BenchmarkMinSubArrayLenHint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minSubArrayLenHint(7, []int{2, 3, 1, 2, 4, 3})
	}
}

// 19.5 ns/op
func BenchmarkMinSubArrayLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	}
}
