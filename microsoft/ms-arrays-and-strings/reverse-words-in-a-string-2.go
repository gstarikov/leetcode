package ms_arrays_and_strings

import (
	"reflect"
	"unsafe"
)

/*
Given a character array s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by a single space.
Your code must solve the problem in-place, i.e. without allocating extra space.
*/

func ReverseWords2(s []byte) {
	// reverse whole string
	reverseString(s)
	// reverse word in reversed string
	i, l := 0, 0
	ln := len(s)
	for i >= 0 {
		//skip spaces
		for ; i < ln && s[i] == ' '; i++ {
		}
		if i >= ln {
			break
		}
		//search word begin
		for l = i; l < ln && s[l] != ' '; l++ {
		}
		// between i & l are word
		reverseString(s[i:l])
		i = l
	}
}

func reverseString(s []byte) {
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
