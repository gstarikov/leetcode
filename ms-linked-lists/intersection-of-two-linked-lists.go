package ms_linked_lists

/*
Given the heads of two singly linked-lists headA and headB,
return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

Constraints:

    The number of nodes of listA is in the m.
    The number of nodes of listB is in the n.
    1 <= m, n <= 3 * 10^4
    1 <= Node.val <= 10^5
    0 <= skipA < m
    0 <= skipB < n
    intersectVal is 0 if listA and listB do not intersect.
    intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.

Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?
*/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	const visitMark = 1e6

	origA := headA
	// run throuht first list and mark
	for ; headA != nil; headA = headA.Next {
		headA.Val += visitMark
	}
	// run throuth second list and find first mark
	for ; headB != nil; headB = headB.Next {
		if headB.Val >= visitMark {
			break
		}
	}
	//clean up nodes from mark
	for headA = origA; headA != nil; headA = headA.Next {
		headA.Val -= visitMark
	}

	return headB
}

func GetIntersectionNodeTwoPointers(headA, headB *ListNode) *ListNode {
	l1 := listLen(headA)
	l2 := listLen(headB)
	longest := headA
	shorter := headB
	ln := l1 - l2
	if l2 > l1 {
		longest = headB
		shorter = headA
		ln = l2 - l1
	}
	for i := ln; i > 0; i-- {
		longest = longest.Next
	}
	for ; longest != shorter; longest, shorter = longest.Next, shorter.Next {
	} //risk of going out of memory border
	return shorter
}
