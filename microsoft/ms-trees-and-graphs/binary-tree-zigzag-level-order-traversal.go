package ms_trees_and_graphs

/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
(i.e., from left to right, then right to left for the next level and alternate between).
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	helperZigZapLevelOrder(root, &ret, 0)
	return ret
}

func helperZigZapLevelOrder(root *TreeNode, data *[][]int, level int) {
	switch {
	case root == nil:
		return
	case len(*data) == level:
		*data = append(*data, []int{})
	case len(*data) < level:
		panic("impossible case")
	}

	val := (*data)[level]
	switch level % 2 {
	case 1:
		val = append([]int{root.Val}, val...)
	case 0:
		val = append(val, root.Val)
	}
	(*data)[level] = val

	helperZigZapLevelOrder(root.Left, data, level+1)
	helperZigZapLevelOrder(root.Right, data, level+1)
}
