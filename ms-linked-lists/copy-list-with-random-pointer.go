package ms_linked_lists

/*
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.
Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.
For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.
Return the head of the copied linked list.

Constraints:
    0 <= n <= 1000
    -104 <= Node.val <= 104
    Node.random is null or is pointing to some node in the linked list.

Constraints:
    0 <= n <= 1000
    -104 <= Node.val <= 104
    Node.random is null or is pointing to some node in the linked list.

*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//do copy
	//interwave copyed nodes
	p := head
	for p != nil {
		nw := &Node{
			Val:    p.Val,
			Next:   p.Next,
			Random: nil,
		}
		p.Next = nw
		p = nw.Next
	}
	//resolve random pointers
	p = head
	for ; p != nil; p = p.Next.Next { //jump over copy node
		//we are on original node
		rOrig := p.Random //got original rnd node
		if rOrig != nil {
			rCpy := rOrig.Next   //got corresponding copy node
			p.Next.Random = rCpy //set on copy valid rnd ptr
		}
	}
	//split lists
	ret := head.Next
	pr := ret
	p = head
	for p != nil {
		nxt := p.Next.Next
		p.Next = nxt
		p = nxt
		if pr.Next != nil {
			cNxt := pr.Next.Next
			pr.Next = cNxt
			pr = cNxt
		}
	}
	return ret
}
