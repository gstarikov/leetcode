package ms_linked_lists

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Constraints:

    The number of nodes in both lists is in the range [0, 50].
    -100 <= Node.val <= 100
    Both list1 and list2 are sorted in non-decreasing order.
*/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret *ListNode
	//chose head
	switch {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	case list1.Val > list2.Val:
		ret = list2
		list2 = list2.Next
	default:
		ret = list1
		list1 = list1.Next
	}

	//run over lists and relink it
	ptr := ret
	for list1 != nil && list2 != nil {
		switch {
		case list1.Val < list2.Val:
			ptr.Next = list1
			ptr = list1
			list1 = list1.Next
		default:
			ptr.Next = list2
			ptr = list2
			list2 = list2.Next
		}
	}

	//attach remaining tail
	switch {
	case list1 == nil:
		ptr.Next = list2
	default:
		ptr.Next = list1
	}
	return ret
}
