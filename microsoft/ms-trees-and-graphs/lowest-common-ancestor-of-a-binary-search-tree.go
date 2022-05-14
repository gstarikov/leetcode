package ms_trees_and_graphs

/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.
According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q
as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Constraints:
    The number of nodes in the tree is in the range [2, 10^5].
    -10^9 <= Node.val <= 10^9
    All Node.val are unique.
    p != q
    p and q will exist in the BST.

*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pv := p.Val
	qv := q.Val

	for root != nil {
		rv := root.Val
		switch {
		case pv > rv && qv > rv:
			root = root.Right
		case pv < rv && qv < rv:
			root = root.Left
		default: //split point
			return root
		}
	}
	return root
}
