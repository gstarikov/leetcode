package ms_arrays_and_strings

import (
	"reflect"
	"unsafe"
)

/*
Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.
*/

func ReverseString(s []byte) {
	p1 := reflect.ValueOf(&s[0]).Pointer()
	p2 := reflect.ValueOf(&s[len(s)-1]).Pointer()
	for p1 < p2 {
		tmp := *(*byte)(unsafe.Pointer(p1))
		*(*byte)(unsafe.Pointer(p1)) = *(*byte)(unsafe.Pointer(p2))
		*(*byte)(unsafe.Pointer(p2)) = tmp
		p1++
		p2--
	}
}
