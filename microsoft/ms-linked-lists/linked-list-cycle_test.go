package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		gen func() *ListNode
		res bool
	}{
		{gen: func() *ListNode {
			var arr []ListNode
			arr = []ListNode{
				{Val: 3},
				{Val: 2},
				{Val: 0},
				{Val: -4},
			}
			arr[0].Next = &arr[1]
			arr[1].Next = &arr[2]
			arr[2].Next = &arr[3]
			arr[3].Next = &arr[1]
			return &arr[0]
		}, res: true},
		{gen: func() *ListNode {
			var arr []ListNode
			arr = []ListNode{
				{Val: 1},
				{Val: 2},
			}
			arr[0].Next = &arr[1]
			arr[1].Next = &arr[0]
			return &arr[0]
		}, res: true},
		{gen: func() *ListNode {
			var arr []ListNode
			arr = []ListNode{
				{Val: 1},
			}
			return &arr[0]
		}, res: false},
	}

	for i, test := range tests {
		in := test.gen()
		r := HasCycle(in)
		assert.Equalf(t, test.res, r, "case(%d)", i)
	}
}
