package array_and_string

import (
	"reflect"
	"unsafe"
)

/*
Given a string s, reverse the order of characters in each word within a sentence
while still preserving whitespace and initial word order.
*/

func reverseWords(s string) string {
	ln := len(s)
	in := []byte(s)
	ret := make([]byte, ln)
	var start, end int
	for {
		start = end + helperFirstNonSpace(in[end:])
		//set spaces - there is no way to use memset :( so use copy
		copy(ret[end:], in[end:start])
		if start >= ln {
			break
		}
		end = start + helperFirstSpace(in[start:])
		helperReverse(in[start:end], ret[start:])
		if end >= ln {
			break
		}
	}
	return string(ret)
}

func helperFirstNonSpace(s []byte) int {
	var i int
	for i = 0; i < len(s) && s[i] == ' '; i++ {
	}
	return i
}
func helperFirstSpace(s []byte) int {
	var i int
	for i = 0; i < len(s) && s[i] != ' '; i++ {
	}
	return i
}

//ms_arrays_and_strings.ReverseString()
func helperReverse(s, t []byte) {
	p1 := reflect.ValueOf(&s[len(s)-1]).Pointer()
	p2 := reflect.ValueOf(&t[0]).Pointer()
	for cnt := len(s); cnt > 0; cnt-- {
		*(*byte)(unsafe.Pointer(p2)) = *(*byte)(unsafe.Pointer(p1))
		p1--
		p2++
	}
}
