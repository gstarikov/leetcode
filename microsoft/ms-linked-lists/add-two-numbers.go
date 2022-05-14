package ms_linked_lists

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}
	var resHead = &ListNode{}
	var resPos = resHead
	var carryFlag int
	// do summing until one of list will end
	for ; (l1 != nil) && (l2 != nil); l1, l2 = l1.Next, l2.Next {
		resPos.Next = new(ListNode)
		resPos = resPos.Next

		var tmpAcc int
		tmpAcc += carryFlag
		tmpAcc += l1.Val
		tmpAcc += l2.Val
		tmpAcc, carryFlag = tmpAcc%10, tmpAcc/10
		resPos.Val = tmpAcc
	}

	//one of lists is empty, do summing last one
	leastList := l1
	if leastList == nil {
		leastList = l2
	}
	for ; leastList != nil; leastList = leastList.Next {
		resPos.Next = new(ListNode)
		resPos = resPos.Next

		var tmpAcc int
		tmpAcc += carryFlag
		tmpAcc += leastList.Val
		tmpAcc, carryFlag = tmpAcc%10, tmpAcc/10
		resPos.Val = tmpAcc
	}

	if carryFlag != 0 {
		resPos.Next = new(ListNode)
		resPos = resPos.Next
		resPos.Val = carryFlag
	}

	resHead = resHead.Next // remove first empty elem
	return resHead
}
