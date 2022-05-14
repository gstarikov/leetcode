package ms_trees_and_graphs

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Constraints:
    The number of nodes in the tree is in the range [0, 2000].
    -1000 <= Node.val <= 1000

*/
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	helperLevelOrder(root, &ret, 0)
	return ret
}

func helperLevelOrder(root *TreeNode, data *[][]int, level int) {
	switch {
	case root == nil:
		return
	case len(*data) == level:
		*data = append(*data, []int{})
	case len(*data) < level:
		panic("impossible case")
	}
	vec := (*data)[level]
	vec = append(vec, root.Val)
	(*data)[level] = vec

	helperLevelOrder(root.Left, data, level+1)
	helperLevelOrder(root.Right, data, level+1)
}
