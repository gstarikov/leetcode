package array_and_string

import "unsafe"

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.
*/

//16 ns/op
func rotateTrivial(nums []int, k int) {
	n1 := len(nums) - 1
	for ; k > 0; k-- {
		prev := nums[n1]
		for i := 0; i <= n1; i++ {
			tmp := nums[i]
			nums[i] = prev
			prev = tmp
		}
	}
}

//speedup copying
//18 ns/op
func rotateTrivial2(nums []int, k int) {
	n1 := len(nums) - 1
	for ; k > 0; k-- {
		prev := nums[n1]
		copy(nums[1:], nums[:n1])
		nums[0] = prev
	}
}

//use additional memory
// 31 ns/op (allocations sucks)
func rotateTrivial3(nums []int, k int) {
	tmp := make([]int, k)
	pos := len(nums) - k
	copy(tmp, nums[pos:]) //save tail
	copy(nums[k:], nums)  //
	copy(nums, tmp)
}

// 71 ns/op :(
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 || n == 0 {
		return
	}

	elems := helperGCDEuclidean(k, n)

	for i := 0; i < elems; i++ {
		prev := nums[i] //slice bound check too costly...
		for pos := i + k; pos != i; pos = (pos + k) % n {
			tmp := nums[pos]
			nums[pos] = prev
			prev = tmp
		}
		nums[i] = prev
	}
}

// 70 ns/op - remove slice bound check
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 || n == 0 {
		return
	}

	elems := helperGCDEuclidean(k, n)

	ptrArr := uintptr(unsafe.Pointer(&nums[0]))

	for i := 0; i < elems; i++ {
		//prev := nums[i]//slice bound check too costly...
		prev := *(*int)(unsafe.Pointer(ptrArr + unsafe.Sizeof(int(0))*uintptr(i))) //nums[i]
		for pos := i + k; pos != i; pos = (pos + k) % n {
			//tmp := nums[pos]
			tmp := *(*int)(unsafe.Pointer(ptrArr + unsafe.Sizeof(int(0))*uintptr(pos))) //nums[pos]
			//nums[pos] = prev
			*(*int)(unsafe.Pointer(ptrArr + unsafe.Sizeof(int(0))*uintptr(pos))) = prev
			prev = tmp
		}
		//nums[i] = prev
		*(*int)(unsafe.Pointer(ptrArr + unsafe.Sizeof(int(0))*uintptr(i))) = prev
	}
}

func helperGCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
