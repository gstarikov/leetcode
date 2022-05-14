package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "hello",
			out: "olleh",
		},
		{
			in:  "alloha",
			out: "aholla",
		},
	}
	for i, test := range tests {
		b := []byte(test.in)
		ReverseString(b)
		assert.Equalf(t, test.out, string(b), "case(%d)", i)
	}
}
