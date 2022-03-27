package backtraking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWildcardMatching(t *testing.T) {
	tests := []struct {
		s, p string
		res  bool
	}{
		{
			s:   "aa",
			p:   "a",
			res: false,
		},
		{
			s:   "aa",
			p:   "*",
			res: true,
		},
		{
			s:   "cb",
			p:   "?a",
			res: false,
		},
		{
			s:   "abaabcab",
			p:   "ab*bc*b",
			res: true,
		},
	}

	for i, test := range tests {
		res := WildcardMatching(test.s, test.p)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
