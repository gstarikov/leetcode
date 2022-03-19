package ms_trees_and_graphs

/*
Given two integer arrays preorder and inorder
where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree,
construct and return the binary tree.

Constraints:
    1 <= preorder.length <= 3000
    inorder.length == preorder.length
    -3000 <= preorder[i], inorder[i] <= 3000
    preorder and inorder consist of unique values.
    Each value of inorder also appears in preorder.
    preorder is guaranteed to be the preorder traversal of the tree.
    inorder is guaranteed to be the inorder traversal of the tree.

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Input: preorder = [-1], inorder = [-1]
Output: [-1]

*/

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	//hashmap to speedup search
	searchMap := map[int]int{}
	for i, v := range inorder {
		searchMap[v] = i
	}

	btc := buildTreeClass{
		search:   searchMap,
		preorder: preorder,
		inorder:  inorder,
	}

	return btc.helperBuildSubtree(0, len(inorder)-1)
}

type buildTreeClass struct {
	search            map[int]int
	preorder, inorder []int
}

func (t *buildTreeClass) helperBuildSubtree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	root := t.preorder[0]
	t.preorder = t.preorder[1:]
	rootIdx := t.search[root]
	return &TreeNode{
		Val:   root,
		Left:  t.helperBuildSubtree(left, rootIdx-1),
		Right: t.helperBuildSubtree(rootIdx+1, right),
	}
}
