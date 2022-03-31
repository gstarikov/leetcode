package ms_2020_03

type List struct {
	key int
	nxt *List
}

func MergeSort(lst *List) *List {
	if lst == nil {
		return nil
	}

	listLen := helperCalcListLen(lst)
	switch listLen {
	case 1:
		return lst
	case 2:
		e1, e2 := lst, lst.nxt
		if e1.key <= e2.key {
			return lst
		} else {
			e2.nxt = e1
			e1.nxt = nil
			return e2
		}
	}
	splitPoint := listLen >> 1
	p1, p2 := helperSpilt(lst, splitPoint)
	s1 := MergeSort(p1)
	s2 := MergeSort(p2)
	res := Merge(s1, s2)
	return res
}

func Merge(l1, l2 *List) *List {
	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var res *List
	if l1.key < l2.key {
		res = l1
		l1 = l1.nxt
	} else {
		res = l2
		l2 = l2.nxt
	}
	head := res

	for l1 != nil && l2 != nil {
		if l1.key <= l2.key {
			res.nxt = l1
			l1 = l1.nxt
		} else {
			res.nxt = l2
			l2 = l2.nxt
		}
		res = res.nxt
	}
	//one of l1 || l2 == nil
	if l1 == nil {
		res.nxt = l2
	} else {
		res.nxt = l1
	}
	return head
}

func helperCalcListLen(lst *List) int {
	var i int
	for i = 0; lst != nil; i, lst = i+1, lst.nxt {
	}
	return i
}

func helperSpilt(lst *List, point int) (p1, p2 *List) {
	p1 = lst
	point--
	var i int
	for p2 = lst; p2 != nil && i < point; i, p2 = i+1, p2.nxt {
	}
	tmp := p2.nxt
	p2.nxt = nil
	p2 = tmp

	return
}
