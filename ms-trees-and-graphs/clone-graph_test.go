package ms_trees_and_graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	res := CloneGraph(nil)
	var nilNode *GraphNode
	assert.Equal(t, nilNode, res)
}
