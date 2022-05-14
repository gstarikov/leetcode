package ms_linked_lists

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		in []int
	}{
		{in: []int{1, 2, 3, 4, 5}},
		{in: []int{1, 2}},
	}
	for _, test := range tests {
		ls := makeList(test.in)
		printList(t, "in", ls)
		out := ReverseList(ls)
		printList(t, "out", out)

	}
}

func makeList(in []int) *ListNode {
	if len(in) == 0 {
		return nil
	}
	var prev *ListNode
	prev = &ListNode{
		Val:  in[0],
		Next: nil,
	}
	head := prev
	for i := 1; i < len(in); i++ {
		tmp := new(ListNode)
		tmp.Val = in[i]
		prev.Next = tmp
		prev = tmp
	}
	return head
}

func printList(t testing.TB, s string, ls *ListNode) {
	for i := 1; ls != nil; ls, i = ls.Next, i+1 {
		t.Logf("%s\t%d->%d", s, i, ls.Val)
	}
}
