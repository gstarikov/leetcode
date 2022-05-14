package ms_linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	n4 := &Node{
		Val:  1,
		Next: nil,
	}
	n3 := &Node{
		Val:  10,
		Next: n4,
	}
	n2 := &Node{
		Val:  11,
		Next: n3,
	}
	n1 := &Node{
		Val:  13,
		Next: n2,
	}
	n0 := &Node{
		Val:  7,
		Next: n1,
	}

	n1.Random = n0
	n2.Random = n4
	n3.Random = n2
	n4.Random = n0

	res := copyRandomList(n0)
	assert.Equal(t, n0, res)
}
