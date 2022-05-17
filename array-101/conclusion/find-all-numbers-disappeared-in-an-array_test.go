package conclusion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			in:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			out: []int{5, 6},
		},
	}
	for i, test := range tests {
		out := findDisappearedNumbers(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
