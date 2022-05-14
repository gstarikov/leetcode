package ms_trees_and_graphs

import "math"

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

    The left subtree of a node contains only nodes with keys less than the node's key.
    The right subtree of a node contains only nodes with keys greater than the node's key.
    Both the left and right subtrees must also be binary search trees.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return checkBST(root, math.MinInt64, math.MaxInt64)
}

func checkBST(root *TreeNode, min, max int) bool {
	v := root.Val
	ret := true
	if root.Left != nil {
		lv := root.Left.Val
		ret = ret && lv < v && lv > min && checkBST(root.Left, min, v)
	}
	if ret && root.Right != nil {
		rv := root.Right.Val
		ret = ret && rv > v && rv < max && checkBST(root.Right, v, max)
	}
	return ret
}
