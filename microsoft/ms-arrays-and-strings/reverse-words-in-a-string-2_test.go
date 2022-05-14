package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords2(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "the sky is blue",
			out: "blue is sky the",
		},
		{
			in:  "a",
			out: "a",
		},
	}
	for i, test := range tests {
		in := []byte(test.in)
		ReverseWords2(in)
		assert.Equalf(t, test.out, string(in), "case(%d)", i)
	}
}
