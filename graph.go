package main

// Node - basic graph node with path cost
type Node struct {
	PathCost int
	Clause   Atom
	Truth    bool
}

// Edge - directed graph edge with weight if graph will be
// weighted
type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

// Graph - the data structure that holds the nodes and their
// connections
type Graph struct {
	Nodes      []*Node
	Edges      map[*Node][]*Node
	TruthNode  *Node
	TypeExists map[string]bool
}

// CreateGraph - creates the initial graph with the truth node as
// the starting node
func CreateGraph() *Graph {
	truthNode := Node{Truth: true}
	g := Graph{Edges: make(map[*Node][]*Node)}
	g.AddNode(&truthNode, nil, 0)
	return &g
}

// AddNode add a node to the graph with an edge weight
func (g *Graph) AddNode(n *Node, from *Node, weight int) {

	// if this is not the first node in the graph
	if from != nil {
		n.PathCost = from.PathCost + weight
		g.Edges[from] = append(g.Edges[from], n)
	} else {
		// first node added to graph, then path cost must be 0
		n.PathCost = 0
	}

	if n.Truth {
		g.TruthNode = n
	}

	g.Nodes = append(g.Nodes, n)

}

// GetTruthNode - returns the truth node in the graph
func (g *Graph) GetTruthNode() *Node {
	return g.TruthNode
}

// GetConnectedNodes - get a list of all the reachable nodes (with distance 1)
// from a source node n
func (g *Graph) GetConnectedNodes(n *Node) (nodes []*Node) {
	return g.Edges[n]
}
