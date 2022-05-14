package inserting_items_into_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckN(t *testing.T) {
	tests := []struct {
		n  []int
		ok bool
	}{
		{
			n:  []int{10, 2, 5, 3},
			ok: true,
		},
		{
			n:  []int{-2, 0, 10, -19, 4, 6, -8},
			ok: false,
		},
		{
			n:  []int{0, 0},
			ok: true,
		},
	}

	for i, test := range tests {
		r := checkIfExist(test.n)
		assert.Equalf(t, test.ok, r, "case(%d)", i)
	}
}
