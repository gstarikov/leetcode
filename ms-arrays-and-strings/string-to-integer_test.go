package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{in: "42", out: 42},
		{in: "   -42", out: -42},
		{in: "4193 with words", out: 4193},
		{in: "", out: 0},
	}

	for i, test := range tests {
		r := MyAtoi(test.in)
		assert.Equalf(t, test.out, r, "case(%d)", i)
	}
}
