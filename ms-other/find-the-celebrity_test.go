package ms_other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		graph [][]int
		out   int
	}{
		{
			graph: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			out: 1,
		},
		{
			graph: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			out: -1,
		},
		{
			graph: [][]int{
				{1, 1},
				{1, 1},
			},
			out: -1,
		},
	}

	for i, test := range tests {
		f := solutionMy(func(a int, b int) bool {
			return test.graph[a][b] == 1
		})
		out := f(len(test.graph))
		assert.Equalf(t, test.out, out, "case(%d)", i)
	}
}
