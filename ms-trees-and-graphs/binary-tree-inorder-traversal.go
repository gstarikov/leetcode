package ms_trees_and_graphs

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

func InorderTraversal(root *TreeNode) []int {
	// moris traversal algoritm
	tourist := root
	var ret []int
	for tourist != nil {
		if tourist.Left == nil {
			ret = append(ret, tourist.Val)
			tourist = tourist.Right
		} else {
			guide := tourist.Left
			//run until end or tourist meet
			for ; guide.Right != nil && guide.Right != tourist; guide = guide.Right {
			}
			if guide.Right == nil { //make a bridge and move tourist left
				guide.Right = tourist
				tourist = tourist.Left
			} else { //break bridge and move tourist right
				guide.Right = nil
				ret = append(ret, tourist.Val)
				tourist = tourist.Right
			}
		}
	}
	return ret
}
