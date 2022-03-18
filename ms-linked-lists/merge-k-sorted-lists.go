package ms_linked_lists

import "math"

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.

Constraints:
    k == lists.length
    0 <= k <= 104
    0 <= lists[i].length <= 500
    -104 <= lists[i][j] <= 104
    lists[i] is sorted in ascending order.
    The sum of lists[i].length will not exceed 104.
*/

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//chose head
	ret := chooseMin(lists)

	//run over lists and relink it
	for ptr := ret; ptr != nil; {
		//scan for min elem
		tmp := chooseMin(lists)
		ptr.Next = tmp
		ptr = tmp
	}

	return ret
}

func chooseMin(lists []*ListNode) *ListNode {
	min := math.MaxInt64
	var tmpHead **ListNode
	for i := range lists {
		elem := &lists[i]
		if *elem != nil && (*elem).Val <= min {
			min = (*elem).Val
			tmpHead = elem
		}
	}
	if tmpHead == nil { //all are nil
		return nil
	}
	ret := *tmpHead            //i have found elem, so it cant be nil
	*tmpHead = (*tmpHead).Next //remove added elem
	return ret
}

// unsuccesfull try to speedup
func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//chose head
	ret, last := chooseMin2(lists)

	//run over lists and relink it
	for ptr := last; ptr != nil; {
		//scan for min elem
		var tmp *ListNode
		tmp, last = chooseMin2(lists)
		ptr.Next = tmp
		ptr = last
	}

	return ret
}

func chooseMin2(lists []*ListNode) (first, last *ListNode) {
	// надо получить минимальный из больших
	minFromGreater := math.MaxInt64
	min := math.MaxInt64
	var tmpHead **ListNode
	for i := range lists {
		elem := &lists[i]
		if *elem == nil {
			continue
		}
		v := (*elem).Val
		switch {
		case v < min:
			minFromGreater = min
			min = (*elem).Val
			tmpHead = elem
		case v == min: //there is two same val chains
			minFromGreater = min
		case v > min && v < minFromGreater:
			minFromGreater = v
		}
	}
	if tmpHead == nil { //all are nil
		return nil, nil
	}
	first = *tmpHead           //i have found elem, so it cant be nil
	*tmpHead = (*tmpHead).Next //remove added elem
	last = first

	// add elems until they are less that. it is speed optimization to avoid useless min scanning in case of subsequence
	tmpPtr := *tmpHead
	for ; tmpPtr != nil && tmpPtr.Val <= minFromGreater; tmpPtr = tmpPtr.Next {
		last = tmpPtr
	}
	*tmpHead = last.Next

	return first, last
}
