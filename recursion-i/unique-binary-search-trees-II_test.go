package recursion_i

import "testing"

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n int
	}{
		{n: 3},
	}

	for i, test := range tests {
		_ = generateTrees(test.n)
		_ = i
	}
}
