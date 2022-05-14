package inserting_items_into_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidMontaun(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{in: []int{2, 1}, out: false},
		{in: []int{3, 5, 5}, out: false},
		{in: []int{0, 3, 2, 1}, out: true},
	}
	for i, test := range tests {
		out := validMountainArray(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
