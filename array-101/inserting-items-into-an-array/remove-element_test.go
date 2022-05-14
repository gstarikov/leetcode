package inserting_items_into_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		in, out []int
		k       int
	}{
		{
			in:  []int{3, 2, 2, 3},
			out: []int{2, 2},
			k:   3,
		},
	}

	for i, test := range tests {
		out := removeElement(test.in, test.k)
		assert.Equalf(t, test.out, test.in[:out], "case(%d)", i)
	}
}
