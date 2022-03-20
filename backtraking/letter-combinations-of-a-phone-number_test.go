package backtraking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "23",
			out: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			in:  "",
			out: []string{},
		},
		{
			in:  "2",
			out: []string{"a", "b", "c"},
		},
	}

	for i, test := range tests {
		res := letterCombinations(test.in)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
