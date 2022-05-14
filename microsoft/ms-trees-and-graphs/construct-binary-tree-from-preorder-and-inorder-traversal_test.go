package ms_trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	nodes := [5]TreeNode{}
	nodes[0] = TreeNode{
		Val:   3,
		Left:  &nodes[1],
		Right: &nodes[2],
	}
	nodes[1] = TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	nodes[2] = TreeNode{
		Val:   20,
		Left:  &nodes[3],
		Right: &nodes[4],
	}
	nodes[3] = TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	nodes[4] = TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	res := BuildTree(preorder, inorder)
	assert.Equal(t, &nodes[0], res)
}
