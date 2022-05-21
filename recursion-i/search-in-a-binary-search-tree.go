package recursion_i

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		switch {
		case val == root.Val:
			return root
		case val < root.Val:
			root = root.Left
		case val > root.Val:
			root = root.Right
		default:
			panic("impossible case")
		}
	}
	return root
}
