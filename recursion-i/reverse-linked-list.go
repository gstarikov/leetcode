package recursion_i

func reverseListIRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, _ := reverseListRec(head, head.Next)
	return p1
}

func reverseListRec(prev, next *ListNode) (head, last *ListNode) {
	if next.Next == nil {
		next.Next = prev
		prev.Next = nil
		return next, prev
	}
	head, last = reverseListRec(next, next.Next)
	last.Next = prev
	prev.Next = nil
	return head, prev
}

func reverseListIter(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
