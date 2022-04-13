package ms_sorting_and_searching

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

func findMedianSortedArraysHint(nums1 []int, nums2 []int) float64 {
	//https://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/

}
