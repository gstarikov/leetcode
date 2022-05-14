package ms_trees_and_graphs

/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”
*/

func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	n, r := searchSubtree2(root, p, q)
	if n == 2 {
		return r
	}
	return nil
}

func searchSubtree2(root, p, q *TreeNode) (cnt int, n *TreeNode) {
	switch { //hit node
	case root == nil:
		return 0, nil
	case root.Val == p.Val:
		ok := searchSubtree1(root.Left, q) || searchSubtree1(root.Right, q)
		if ok {
			return 2, root
		}
		return 1, root
	case root.Val == q.Val:
		ok := searchSubtree1(root.Left, p) || searchSubtree1(root.Right, p)
		if ok {
			return 2, root
		}
		return 1, root
	}
	lCnt, lR := searchSubtree2(root.Left, p, q)
	if lCnt == 2 {
		return 2, lR
	}
	rCnt, rR := searchSubtree2(root.Right, p, q)
	if rCnt == 2 {
		return 2, rR
	}
	if lCnt == 1 && rCnt == 1 { //we are split point
		return 2, root
	}
	return lCnt + rCnt, nil
}

func searchSubtree1(root, p *TreeNode) (ok bool) {
	if root == nil {
		return false
	}
	if root.Val == p.Val {
		return true
	}

	return searchSubtree1(root.Left, p) || searchSubtree1(root.Right, p)
}
