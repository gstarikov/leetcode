package array_and_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b, res string
	}{
		{
			a:   "11",
			b:   "1",
			res: "100",
		},
		{
			a:   "1010",
			b:   "1011",
			res: "10101",
		},
	}

	for i, test := range tests {
		res := addBinary(test.a, test.b)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
