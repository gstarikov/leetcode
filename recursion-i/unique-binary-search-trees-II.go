package recursion_i

/*
Given an integer n, return all the structurally unique BST's (binary search trees),
which has exactly n nodes of unique values from 1 to n. Return the answer in any order.
*/

func generateTrees(n int) []*TreeNode {

	return generateSubTree(1, n)
}

func generateSubTree(f, l int) []*TreeNode {
	var ret []*TreeNode

	for i := f; i <= l; i++ {
		lSub := generateSubTree(f, i-1)
		rSub := generateSubTree(i+1, l)
		switch {
		case len(lSub) == 0 && len(rSub) == 0:
			ret = append(ret, &TreeNode{
				Val:   i,
				Left:  nil,
				Right: nil,
			})
		case len(lSub) == 0:
			for _, r := range rSub {
				ret = append(ret, &TreeNode{
					Val:   i,
					Left:  nil,
					Right: r,
				})
			}
		case len(rSub) == 0:
			for _, r := range lSub {
				ret = append(ret, &TreeNode{
					Val:   i,
					Left:  r,
					Right: nil,
				})
			}
		default:
			for _, lr := range lSub {
				for _, rr := range rSub {
					ret = append(ret, &TreeNode{
						Val:   i,
						Left:  lr,
						Right: rr,
					})
				}
			}
		}
	}

	return ret
}
