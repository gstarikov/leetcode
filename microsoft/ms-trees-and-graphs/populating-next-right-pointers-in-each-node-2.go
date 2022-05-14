package ms_trees_and_graphs

func Connect2(root *Node) *Node {
	goDeep(root)
	return root
}

func goDeep(root *Node) {
	if root == nil {
		return
	}
	wireLevel2(root)
	goDeep(root.Left)
	goDeep(root.Right)
}

func wireLevel2(prev *Node) {
	if prev == nil ||
		(prev.Left != nil && prev.Left.Next != nil) || //already visited
		(prev.Right != nil && prev.Right.Next != nil) { //already visited
		return
	}

	l := &Node{} //so dummy
	for ; prev != nil; prev = prev.Next {
		if prev.Left != nil {
			l.Next = prev.Left
			l = prev.Left
		}
		if prev.Right != nil {
			l.Next = prev.Right
			l = prev.Right
		}
	}
}
