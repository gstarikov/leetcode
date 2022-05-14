package ms_arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s   string
		res bool
	}{
		{
			s:   "A man, a plan, a canal: Panama",
			res: true,
		},
		{
			s:   "race a car",
			res: false,
		},
		{
			s:   " ",
			res: true,
		},
	}

	for i, test := range tests {
		res := IsPalindrome(test.s)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
