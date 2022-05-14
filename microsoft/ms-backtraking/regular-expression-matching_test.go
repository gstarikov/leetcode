package ms_backtraking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		out  bool
	}{
		{s: "aa", p: "a", out: false},                   //0
		{s: "aa", p: "a*", out: true},                   //1
		{s: "ab", p: ".*", out: true},                   //2
		{s: "mississippi", p: "mis*is*ip*.", out: true}, //3
		{s: "ab", p: "*.c", out: false},                 //4
		{s: "ab", p: ".*c", out: false},                 //5
		{s: "aaa", p: "ab*a*c*a", out: true},            //6
	}

	for i, test := range tests {
		res := isMatchDP(test.s, test.p)
		assert.Equalf(t, test.out, res, "case(%d)", i)
	}
}
