package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			in:  "Let's take LeetCode contest",
			out: "s'teL ekat edoCteeL tsetnoc",
		}, {
			in:  "God Ding",
			out: "doG gniD",
		},
	}

	for i, test := range tests {
		res := reverseWords(test.in)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
