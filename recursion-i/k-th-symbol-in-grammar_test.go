package recursion_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthGrammar(t *testing.T) {
	tests := []struct {
		n, k, res int
	}{{
		n:   1,
		k:   1,
		res: 0,
	}, {
		n:   2,
		k:   1,
		res: 0,
	}, {
		n:   2,
		k:   2,
		res: 1,
	}, {
		n:   30,
		k:   434991989,
		res: 0,
	}, {
		n:   30,
		k:   434991989,
		res: 0,
	}, {
		n:   4,
		k:   5,
		res: 1,
	}, {
		n:   3,
		k:   3,
		res: 1,
	},
	}

	for i, test := range tests {
		//r := kthGrammar(test.n, test.k)
		r := kthGrammarRec(test.n, test.k)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
