package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "babad",
			out: "bab",
		},
		{
			in:  "cbbd",
			out: "bb",
		},
		{
			in:  "vvv",
			out: "vvv",
		},
		{
			in:  "aaaa",
			out: "aaaa",
		},
	}
	for i, test := range tests {
		res := LongestPalindrome(test.in)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
