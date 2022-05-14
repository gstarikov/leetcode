package ms_trees_and_graphs

/*
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Constraints:

    The number of nodes in the tree is in the range [0, 212 - 1].
    -1000 <= Node.val <= 1000

Follow-up:
    You may only use constant extra space.
    The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	for ptr := root; ptr != nil; ptr = ptr.Left {
		wireLevel(ptr)
	}
	return root
}

func wireLevel(prev *Node) {
	if prev == nil || prev.Left == nil {
		return
	}

	l := &Node{} //so dummy
	for ; prev != nil; prev = prev.Next {
		l.Next = prev.Left
		prev.Left.Next = prev.Right
		l = prev.Right
	}
}
