package ms_sorting_and_searching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		out          float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			out:   2,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			out:   2.5,
		},
	}

	for i, test := range tests {
		res := findMedianSortedArrays(test.nums1, test.nums2)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
