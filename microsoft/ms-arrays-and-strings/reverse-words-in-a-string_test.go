package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "the sky is blue",
			out: "blue is sky the",
		},
		{
			in:  "  hello world  ",
			out: "world hello",
		},
		{
			in:  "a good   example",
			out: "example good a",
		},
		{
			in:  "a",
			out: "a",
		},
	}
	for i, test := range tests {
		r := ReverseWords(test.in)
		assert.Equalf(t, test.out, r, "case(%d)", i)
	}
}
