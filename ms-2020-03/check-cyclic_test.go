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
		{ //0
			arr: []pack{
				{id: "a", dep: []string{"b"}},
				{id: "b", dep: []string{""}},
			},
			res: false,
		},
		{ //1
			arr: []pack{
				{id: "a", dep: []string{"b"}},
				{id: "b", dep: []string{"a"}},
			},
			res: true,
		},
		{ //2
			arr: []pack{
				{id: "a", dep: []string{"b"}},
				{id: "b", dep: []string{"c"}},
				{id: "c", dep: []string{"a"}},
			},
			res: true,
		},
		{ //3
			arr: []pack{
				{id: "a", dep: []string{"b", "c"}},
				{id: "b", dep: []string{}},
				{id: "c", dep: []string{"b"}},
			},
			res: false,
		},
		{ //4
			arr: []pack{
				{id: "a", dep: []string{"b", "c"}},
				{id: "b", dep: []string{"c"}},
				{id: "c", dep: []string{}},
			},
			res: false,
		},
		{ //5
			arr: []pack{
				{id: "a", dep: []string{"b", "c"}}, //test for hangup
				{id: "b", dep: []string{"c"}},
				{id: "c", dep: []string{"b"}},
			},
			res: true,
		},
	}

	for i, test := range tests {
		res := CheckCyc(test.arr)
		assert.Equalf(t, test.res, res, "case(%d)", i)
	}
}
