package ms_trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	nodes := [7]Node{}
	nodes[0] = Node{
		Val:   1,
		Left:  &nodes[1],
		Right: &nodes[2],
		Next:  nil,
	}
	nodes[1] = Node{
		Val:   2,
		Left:  &nodes[3],
		Right: &nodes[4],
		Next:  nil,
	}
	nodes[2] = Node{
		Val:   3,
		Left:  &nodes[5],
		Right: &nodes[6],
		Next:  nil,
	}
	nodes[3] = Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	nodes[4] = Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	nodes[5] = Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	nodes[6] = Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	esp := [7]Node{}
	esp[0] = Node{
		Val:   1,
		Left:  &esp[1],
		Right: &esp[2],
		Next:  nil,
	}
	esp[1] = Node{
		Val:   2,
		Left:  &esp[3],
		Right: &esp[4],
		Next:  &esp[2],
	}
	esp[2] = Node{
		Val:   3,
		Left:  &esp[5],
		Right: &esp[6],
		Next:  nil,
	}
	esp[3] = Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  &esp[4],
	}
	esp[4] = Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  &esp[5],
	}
	esp[5] = Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  &esp[6],
	}
	esp[6] = Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	r := Connect(&nodes[0])
	assert.Equal(t, &esp[0], r)
}
