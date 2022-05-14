package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		in  string
		res bool
	}{
		{
			in:  "()",
			res: true,
		},
		{
			in:  "()[]{}",
			res: true,
		},
		{
			in:  "(]",
			res: false,
		},
		{
			in:  "(",
			res: false,
		},
		{
			in:  ")",
			res: false,
		},
	}

	for i, test := range tests {
		r := ValidParentheses(test.in)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
