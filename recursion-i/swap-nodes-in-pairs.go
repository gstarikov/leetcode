package recursion_i

/*
Given a linked list, swap every two adjacent nodes and return its head.
You must solve the problem without modifying the values in the list's nodes
(i.e., only nodes themselves may be changed.)

Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:

Input: head = []
Output: []

Example 3:

Input: head = [1]
Output: [1]

Constraints:

    The number of nodes in the list is in the range [0, 100].
    0 <= Node.val <= 100
*/

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ret ListNode
	swapPairRec(&ret, head)
	return ret.Next
}

func swapPairRec(prv, nxt *ListNode) *ListNode {
	if nxt == nil || nxt.Next == nil {
		return prv
	}
	p1, p2 := swapPair(nxt)
	prv.Next = p1
	return swapPairRec(p2, p2.Next)
}

func swapPairsNoRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//swap first elems because there is no prev
	ret, prv := swapPair(head)

	for prv != nil && prv.Next != nil && prv.Next.Next != nil {
		p1, p2 := swapPair(prv.Next)
		prv.Next = p1
		prv = p2
	}
	return ret
}

func swapPair(head *ListNode) (p1, p2 *ListNode) {
	nxt := head.Next.Next
	p1, p2 = head, head.Next
	p2.Next = p1
	p1.Next = nxt
	return p2, p1
}
