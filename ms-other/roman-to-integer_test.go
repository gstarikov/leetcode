package ms_other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input string
		res   int
	}{
		{input: "III", res: 3},
		{input: "LVIII", res: 58},
		{input: "MCMXCIV", res: 1994},
		{input: "IV", res: 4},
		{input: "MCDLXXVI", res: 1476},
	}

	for i, test := range tests {
		res := romanToInt(test.input)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}

/*
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
//two maps - common and special cases
BenchmarkRomanTiInt-8   	 4285208	       263.3 ns/op
// replace one map with switch
BenchmarkRomanTiInt-8   	 6013010	       184.7 ns/op
// replace all maps with switch
BenchmarkRomanTiInt-8   	90308852	        12.77 ns/op
*/
func BenchmarkRomanTiInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = romanToInt("MCDLXXVI")
	}
}
