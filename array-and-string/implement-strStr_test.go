package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		str, sub string
		ret      int
	}{
		{
			str: "hello",
			sub: "ll",
			ret: 2,
		},
		{
			str: "aaaaa",
			sub: "ba",
			ret: -1,
		},
	}

	for i, test := range tests {
		ret := strStr(test.str, test.sub)
		assert.Equalf(t, test.ret, ret, "case(%d)", i)
	}
}
