package recursion_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n, cases int
	}{
		{
			n:     0,
			cases: 0,
		},
		{
			n:     1,
			cases: 1,
		},
		{
			n:     2,
			cases: 2,
		},
		{
			n:     3,
			cases: 3,
		},
		{
			n:     4,
			cases: 5,
		},
		{
			n:     5,
			cases: 8,
		},
	}

	for i, test := range tests {
		r := climbStairs(test.n)
		assert.Equalf(t, test.cases, r, "case(%d)", i)
	}
}
