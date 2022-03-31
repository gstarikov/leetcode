package ms_2020_03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckCyc(t *testing.T) {
	tests := []struct {
		arr []pack
		res bool
	}{
		{
			arr: []pack{
				{
					id:  "a",
					dep: []string{"b"},
				},
				{
					id:  "b",
					dep: []string{""},
				},
			},
			res: false,
		},
		{
			arr: []pack{
				{
					id:  "a",
					dep: []string{"b"},
				},
				{
					id:  "b",
					dep: []string{"a"},
				},
			},
			res: true,
		},
		{
			arr: []pack{
				{
					id:  "a",
					dep: []string{"b"},
				},
				{
					id:  "b",
					dep: []string{"c"},
				},
				{
					id:  "c",
					dep: []string{"a"},
				},
			},
			res: true,
		},
	}

	for i, test := range tests {
		res := CheckCyc(test.arr)
		//t.Errorf()
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
