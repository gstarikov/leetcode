package ms_other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		input string
		res   int
	}{
		{input: "A", res: 1},
		{input: "AB", res: 28},
		{input: "ZY", res: 701},
	}

	for i, test := range tests {
		res := titleToNumber(test.input)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
