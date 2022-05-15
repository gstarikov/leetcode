package conclusion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{1, 0, 1, 1, 0},
			out: 4,
		},
		{
			in:  []int{1, 0, 1, 1, 0, 1},
			out: 4,
		},
	}

	for i, test := range tests {
		out := findMaxConsecutiveOnes(test.in)
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
