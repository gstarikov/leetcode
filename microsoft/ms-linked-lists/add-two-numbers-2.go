package ms_linked_lists

/*
You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Constraints:
    The number of nodes in each linked list is in the range [1, 100].
    0 <= Node.val <= 9
    It is guaranteed that the list represents a number that does not have leading zeros.

Follow up: Could you solve it without reversing the input lists?
*/

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	switch { //border cases
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	len1, len2 := listLen(l1), listLen(l2)

	//proceed head
	var res *ListNode
	// len be l1 > l2. so do it
	switch {
	case len1 > len2:
		res, l1 = copyN(l1, len1-len2)
	case len1 < len2:
		tmp := l1
		res, l1 = copyN(l2, len2-len1)
		l2 = tmp
	case len1 == len2:
		//do nothing
	}

	//sum the rest
	for ; l1 != nil; l1, l2 = l1.Next, l2.Next { //only one check because theirs len are equal
		newElem := new(ListNode)
		newElem.Next = res
		res = newElem
		res.Val = l1.Val + l2.Val
	}
	//reverse list & work on carry
	var carry int
	var prev *ListNode
	for res != nil {
		res.Val += carry
		carry = res.Val / 10
		res.Val = res.Val % 10
		nxt := res.Next
		res.Next = prev
		prev = res
		res = nxt

	}
	//last carry
	if carry != 0 {
		prev = &ListNode{
			Val:  carry,
			Next: prev,
		}
	}

	return prev
}

func listLen(l *ListNode) int {
	var ret int
	for ; l != nil; ret, l = ret+1, l.Next {
	}
	return ret
}

func copyN(l *ListNode, n int) (res, head *ListNode) {
	head = l
	for ; head != nil && n > 0; head, n = head.Next, n-1 {
		newElem := new(ListNode)
		newElem.Next = res
		res = newElem
		res.Val = head.Val
	}
	return
}
