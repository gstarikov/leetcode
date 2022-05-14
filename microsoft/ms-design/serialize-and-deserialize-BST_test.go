package ms_design

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(t *testing.T) {
	tests := []struct {
		coded string
	}{
		{coded: "null;"},
		{coded: "1;null;null;"},
		{coded: "1;null;2;null;null;"},
		{coded: "1;null;2;3;null;null;null;"},
	}

	for i, test := range tests {
		codec := BSTConstructor()
		tree := codec.deserialize(test.coded)
		res := codec.serialize(tree)
		assert.Equalf(t, test.coded, res, "case(%d)", i)
	}
}

func TestBSTBin(t *testing.T) {
	tp := &TreeNode{}
	tests := []struct {
		coded string
	}{
		{coded: ""},
		{coded: string(convUint16ToBytes(codeUint16(&TreeNode{Val: 1})))},
		{coded: string(convUint16ToBytes(codeUint16(&TreeNode{Val: 1, Right: tp}))) +
			string(convUint16ToBytes(codeUint16(&TreeNode{Val: 2}))),
		},
		{coded: string(convUint16ToBytes(codeUint16(&TreeNode{Val: 1, Right: tp}))) +
			string(convUint16ToBytes(codeUint16(&TreeNode{Val: 2, Left: tp}))) +
			string(convUint16ToBytes(codeUint16(&TreeNode{Val: 3}))),
		},
		//{coded: "1;null;2;3;null;null;null;"},
	}

	for i, test := range tests {
		codec := BSTConstructor()
		tree := codec.deserialize(test.coded)
		res := codec.serialize(tree)
		assert.Equalf(t, test.coded, res, "case(%d)", i)
	}
}
