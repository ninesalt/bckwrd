package main

// Node - basic graph node with path cost
type Node struct {
	PathCost int
	//TODO add data/clause/atom field
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
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddNode(n *Node, from *Node, weight int) {

	// if this is not the first node in the graph
	if from != nil {
		n.PathCost = from.PathCost + weight
		e := Edge{from, n, weight}
		g.Edges = append(g.Edges, &e)
	} else {
		// first node added to graph, then path cost must be 0
		n.PathCost = 0
	}

	g.Nodes = append(g.Nodes, n)

}
