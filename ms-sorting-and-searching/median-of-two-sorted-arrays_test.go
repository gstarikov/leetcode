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
		{ //0
			nums1: []int{1, 3},
			nums2: []int{2},
			out:   2,
		},
		{ //1
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			out:   2.5,
		},
		{ //2
			nums1: []int{1, 2},
			nums2: []int{-1, 3},
			out:   1.5,
		},
		{ //3
			nums1: []int{1},
			nums2: []int{2, 3, 4},
			out:   2.5,
		},
		{ //4
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6, 7},
			out:   4,
		},
		{ //5
			nums1: []int{2, 2, 4, 4},
			nums2: []int{2, 2, 4, 4},
			out:   3,
		},
		{ //6
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6, 7, 8},
			out:   4.5,
		},
		{ //7
			nums1: []int{1, 4, 7, 10, 12},
			nums2: []int{2, 3, 6, 15},
			out:   6,
		},
	}

	for i, test := range tests {
		//res := findMedianSortedArrays(test.nums1, test.nums2)
		//res := findMedianSortedArraysHint(test.nums1, test.nums2)
		res := findMedianSortedArraysHint2(test.nums1, test.nums2)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}

func TestHelperMedianTwo(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		out          float64
	}{
		{ //0
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			out:   (2 + 3) / 2.0,
		},
		{ //1
			nums1: []int{3, 4},
			nums2: []int{1, 2},
			out:   (2 + 3) / 2.0,
		},
		{ //2
			nums1: []int{1, 4},
			nums2: []int{2, 3},
			out:   (2 + 3) / 2.0,
		},
		{ //3
			nums1: []int{2, 3},
			nums2: []int{1, 4},
			out:   (2 + 3) / 2.0,
		},
		{ //4
			nums1: []int{1, 3},
			nums2: []int{2, 4},
			out:   (2 + 3) / 2.0,
		},
		{ //5
			nums1: []int{2, 4},
			nums2: []int{1, 3},
			out:   (2 + 3) / 2.0,
		},
	}

	for i, test := range tests {
		res := helperMedianOfTwo22(test.nums1, test.nums2)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
