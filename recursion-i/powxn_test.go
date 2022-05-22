package recursion_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		x, res float64
		n      int
	}{
		{ //0
			x:   2,
			res: 4,
			n:   2,
		},
		{ //1
			x:   2,
			res: 8,
			n:   3,
		},
		{ //2
			x:   2,
			res: 16,
			n:   4,
		},
		{ //3
			x:   2,
			res: 32,
			n:   5,
		},
		{ //4
			x:   2,
			res: 1024,
			n:   10,
		},
		{ //5
			x:   8.84372,
			res: 2e-5,
			n:   -5,
		},
	}

	for i, test := range tests {
		r := myPow(test.x, test.n)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
