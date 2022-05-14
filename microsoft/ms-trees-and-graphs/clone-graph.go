package ms_trees_and_graphs

/*
Given a reference of a node in a connected undirected graph.
Return a deep copy (clone) of the graph.
Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

Constraints:
    The number of nodes in the graph is in the range [0, 100].
    1 <= Node.val <= 100
    Node.val is unique for each node.
    There are no repeated edges and no self-loops in the graph.
    The Graph is connected and all nodes can be visited starting from the given node.

*/

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func CloneGraph(node *GraphNode) *GraphNode {
	cgs := cloneGraphStruct{
		alreadyCopyed: map[*GraphNode]*GraphNode{},
	}
	cgs.alreadyCopyed[nil] = nil
	return cgs.clone(node)
}

type cloneGraphStruct struct {
	alreadyCopyed map[*GraphNode]*GraphNode
}

func (p *cloneGraphStruct) clone(node *GraphNode) *GraphNode {
	k, ok := p.alreadyCopyed[node]
	if ok {
		return k
	}
	newNode := &GraphNode{
		Val:       node.Val,
		Neighbors: make([]*GraphNode, len(node.Neighbors)),
	}
	p.alreadyCopyed[node] = newNode

	for i, v := range node.Neighbors {
		newNode.Neighbors[i] = p.clone(v)
	}
	return newNode
}
