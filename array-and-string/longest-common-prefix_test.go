package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		res  string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			res:  "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			res:  "",
		},
		{
			strs: []string{""},
			res:  "",
		},
	}

	for i, test := range tests {
		res := longestCommonPrefix(test.strs)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
