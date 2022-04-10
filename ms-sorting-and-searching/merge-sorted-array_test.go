package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1, nums2, res []int
		m, n              int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			res:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			res:   []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			res:   []int{1},
		},
		{
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
			res:   []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, test := range tests {
		t1 := append([]int{}, test.nums1...)
		Merge(t1, test.m, test.nums2, test.n)
		assert.Equalf(t, test.res, t1, "merge case(%d)", i)

		t2 := append([]int{}, test.nums1...)
		MergeBackward(t2, test.m, test.nums2, test.n)
		assert.Equalf(t, test.res, t2, "mergeBackward case(%d)", i)
	}
}

// report an issue to leetcode - https://github.com/LeetCode-Feedback/LeetCode-Feedback/issues/6963
func BenchmarkMergeBackward(b *testing.B) {
	const ln = 1e4
	t1 := make([]int, ln)
	n := ln >> 1
	m := ln - n
	t2 := make([]int, n)
	for i := 0; i < ln; i++ {
		if i%2 == 0 {
			t1[i/2] = i
		} else {
			t2[i/2] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := append([]int{}, t1...)
		MergeBackward(tmp, n, t2, m)
	}
}
