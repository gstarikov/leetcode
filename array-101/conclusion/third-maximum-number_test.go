package conclusion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThirdMax(t *testing.T) {
	tests := []struct {
		in       []int
		thridmax int
	}{
		{ //0
			in:       []int{3, 2, 1},
			thridmax: 1,
		},
		{ //1
			in:       []int{2, 1},
			thridmax: 2,
		},
		{ //2
			in:       []int{2, 2, 3, 1},
			thridmax: 1,
		},
		{ //3
			in:       []int{1, 2, 2, 5, 3, 5},
			thridmax: 2,
		},
	}
	for i, test := range tests {
		max := thirdMax(test.in)
		assert.Equalf(t, test.thridmax, max, "case(%d)", i)
	}
}
