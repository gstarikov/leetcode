package ms_sorting_and_searching

import (
	"reflect"
	"unsafe"
)

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Constraints:

    nums1.length == m + n
    nums2.length == n
    0 <= m, n <= 200
    1 <= m + n <= 200
    -10^9 <= nums1[i], nums2[j] <= 10^9
Follow up: Can you come up with an algorithm that runs in O(m + n) time?
*/

//cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
//BenchmarkMergeBackward-8            1167           1004994 ns/op
func Merge(nums1 []int, m int, nums2 []int, n int) {
	var idx1, idx2 int
	filled := m
	for idx1 < filled && idx2 < n {
		p1, p2 := &nums1[idx1], &nums2[idx2]
		switch {
		case *p1 >= *p2: //need insert path of p2 into p1
			var i int
			for i = idx2; i < n && *p1 >= nums2[i]; i++ {
			}
			//nums2[idx2:i) -> nums1[idx1]
			ln := i - idx2
			copy(nums1[idx1+ln:], nums1[idx1:filled])
			copy(nums1[idx1:], nums2[idx2:i])
			idx2 += ln
			filled += ln
			idx1 += ln

		case *p1 < *p2: //skip p1 until it is less that p2
			for ; idx1 < filled && nums1[idx1] < *p2; idx1++ {
			}
		}
	}
	//copy tail
	copy(nums1[idx1:], nums2[idx2:n])
}

// it is from solution :( i forgot reverse ordering
//cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
//BenchmarkMergeBackward-8           54319             21828 ns/op
func MergeBackward(nums1 []int, m int, nums2 []int, n int) {
	switch {
	case m == 0:
		copy(nums1, nums2)
		return
	case n == 0:
		return
	}

	step := unsafe.Sizeof(nums1[0]) //be aware of spaces between array elements

	iPtr := reflect.ValueOf(&nums1[m+n-1]).Pointer()
	n1s := reflect.ValueOf(&nums1[0]).Pointer()
	n2s := reflect.ValueOf(&nums2[0]).Pointer()

	p1 := reflect.ValueOf(&nums1[m-1]).Pointer()
	p2 := reflect.ValueOf(&nums2[n-1]).Pointer()

	for p1 >= n1s && p2 >= n2s {
		tp1 := (*int)(unsafe.Pointer(p1))
		tp2 := (*int)(unsafe.Pointer(p2))
		ti := (*int)(unsafe.Pointer(iPtr))

		if *tp1 >= *tp2 {
			*ti = *tp1
			p1 -= step
		} else {
			*ti = *tp2
			p2 -= step
		}
		iPtr -= step
	}
	//copy tail
	if p2 >= n2s {
		//copy n2s:p2 to p1
		//calc indexes
		idx := (p2 - n2s) / step
		copy(nums1, nums2[:idx+1])
	}
}
