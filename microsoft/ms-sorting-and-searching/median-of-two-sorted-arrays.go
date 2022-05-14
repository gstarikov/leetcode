package ms_sorting_and_searching

import "math"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
    nums1.length == m
    nums2.length == n
    0 <= m <= 1000
    0 <= n <= 1000
    1 <= m + n <= 2000
    -10^6 <= nums1[i], nums2[i] <= 10^6
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return helperFindMedianSortedArray(nums1, nums2)
}

// https://takeuforward.org/data-structure/median-of-two-sorted-arrays-of-different-sizes/
// но я тут не вижу бинарного поиска :( тут банально начинаем с середины :(
func findMedianSortedArraysHint2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}
	m, n := len(nums1), len(nums2)
	low, high, medianPos := 0, m, (m+n+1)/2
	for low <= high {
		cut1 := (low + high) >> 1
		cut2 := medianPos - cut1

		l1 := helperTetrary(cut1 == 0, math.MinInt64, cut1-1, nums1)
		l2 := helperTetrary(cut2 == 0, math.MinInt64, cut2-1, nums2)
		r1 := helperTetrary(cut1 == m, math.MaxInt64, cut1, nums1)
		r2 := helperTetrary(cut2 == n, math.MaxInt64, cut2, nums2)

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 != 0 {
				return float64(helperMax(l1, l2))
			}
			return float64(helperMax(l1, l2)+helperMin(r1, r2)) / 2.0
		}
		if l1 > r2 {
			high = cut1 - 1
		} else {
			low = cut1 + 1
		}
	}
	panic("impossible case")
}

func helperMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func helperMin(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func helperTetrary(cond bool, val, pos int, nums []int) int {
	if cond {
		return val
	}
	return nums[pos]

}

//https://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/
// не работает... перепроверил всё, в этом варианте банально кривая разбивка
func findMedianSortedArraysHint(nums1 []int, nums2 []int) float64 {
	for {
		if len(nums1) > len(nums2) {
			tmp := nums1
			nums1 = nums2
			nums2 = tmp
		}

		switch len(nums1) {
		case 0:
			v, _, _ := helperMedianOfOne(nums2)
			return v
		case 1:
			switch {
			case len(nums2) == 1:
				return float64(nums1[0]+nums2[0]) / 2.0
			case len(nums2)%2 == 1:
				mdIdx := len(nums2) / 2
				tmp := nums2[mdIdx-1 : mdIdx+2]
				return helperFindMedianSortedArray(nums1, tmp)
			case len(nums2)%2 == 0:
				mdIdx := len(nums2) / 2
				tmp := nums2[mdIdx-1 : mdIdx+1]
				return helperFindMedianSortedArray(nums1, tmp)
			}
		case 2:
			switch {
			case len(nums2) == 2:
				return helperFindMedianSortedArray(nums1, nums2)
			case len(nums2)%2 == 1:
				mdIdx := len(nums2) / 2
				tmp := nums2[mdIdx-1 : mdIdx+2]
				return helperFindMedianSortedArray(nums1, tmp)
			case len(nums2)%2 == 0:
				mdIdx := len(nums2) / 2
				tmp := nums2[mdIdx-2 : mdIdx+2]
				return helperFindMedianSortedArray(nums1, tmp)
			}
		}

		// do splitting
		m1, idx11, idx12 := helperMedianOfOne(nums1)
		m2, idx21, idx22 := helperMedianOfOne(nums2)

		//decide
		switch {
		case m1 < m2:
			nums1 = nums1[idx11:]
			nums2 = nums2[:idx22+1]
		default:
			nums2 = nums2[idx21:]
			nums1 = nums1[:idx12+1]
		}
	}
}

func helperMedianOfOne(nums []int) (float64, int, int) {
	switch ln := len(nums); ln {
	case 0:
		return -1, 0, 0
	default:
		md1Idx, md2Idx := (ln-1)/2, ln/2
		if md1Idx == md2Idx {
			return float64(nums[md1Idx]), md1Idx, md1Idx
		} else {
			return float64(nums[md1Idx]+nums[md2Idx]) / 2.0, md1Idx, md2Idx
		}
	}
}

func helperFindMedianSortedArray(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total == 0 {
		return 0
	}

	var pos, p1, p2, val1, val2 int
	md1Idx, md2Idx := (total-1)/2, total/2
	p1, p2, pos, val1 = helperPoint(nums1, nums2, p1, p2, pos, md1Idx)
	if md1Idx == md2Idx {
		return float64(val1)
	}
	_, _, _, val2 = helperPoint(nums1, nums2, p1, p2, pos, md2Idx)

	med := float64(val1+val2) / 2
	return med
}

func helperPoint(nums1, nums2 []int, p1, p2, pos int, md int) (int, int, int, int) {
	var lastVal int
	for pos <= md {
		switch {
		case p1 == len(nums1):
			lastVal = nums2[p2] //DRY ?
			p2++
		case p2 == len(nums2):
			lastVal = nums1[p1] //DRY ?
			p1++
		case nums1[p1] < nums2[p2]:
			lastVal = nums1[p1]
			p1++
		case nums1[p1] >= nums2[p2]:
			lastVal = nums2[p2]
			p2++
		}
		pos++
	}
	return p1, p2, pos, lastVal
}

func helperMedianOfTwo12(nums1 int, nums2 []int) float64 {
	if len(nums2) != 2 {
		panic("invalid input data")
	}

	// do quick merge and found median
	switch {
	case nums1 <= nums2[0]:
		return float64(nums2[0])
	case nums1 <= nums2[1]:
		return float64(nums1)
	default:
		return float64(nums2[1])
	}
}

func helperMedianOfTwo13(nums1 int, nums2 []int) float64 {
	if len(nums2) != 3 {
		panic("invalid input data")
	}

	// do quick merge and found median
	switch {
	case nums1 <= nums2[0]:
		return float64(nums2[1]+nums2[0]) / 2.0
	case nums1 <= nums2[2]:
		return float64(nums1+nums2[1]) / 2.0
	case nums1 > nums2[2]:
		return float64(nums2[1]+nums2[2]) / 2.0
	default:
		panic("impossible case")
	}
}

func helperMedianOfTwo22(nums1, nums2 []int) float64 {
	if len(nums1) != 2 || len(nums2) != 2 {
		panic("invalid input data")
	}

	// do quick merge and found median
	var m1, m2 float64
	switch {
	case nums1[1] <= nums2[0]:
		return float64(nums1[1]+nums2[0]) / 2.0
	case nums2[1] <= nums1[0]:
		return float64(nums1[0]+nums2[1]) / 2.0
	case nums1[0] >= nums2[0]:
		m1 = float64(nums1[0])
		if nums1[1] > nums2[1] {
			m2 = float64(nums2[1])
		} else {
			m2 = float64(nums1[1])
		}
	case nums2[0] >= nums1[0]: //DRY !?
		m1 = float64(nums2[0])
		if nums2[1] > nums1[1] {
			m2 = float64(nums1[1])
		} else {
			m2 = float64(nums2[1])
		}
	default:
		panic("impossible case")
	}
	return (m1 + m2) / 2.0
}
