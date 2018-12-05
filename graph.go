package main

// Node - basic graph node with path cost
type Node struct {
	PathCost int
	Clause   *TreeNode
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
	Nodes       []*Node
	TruthNode   *Node
	Edges       map[*Node][]*Node
	TreeNodeMap map[*TreeNode]*Node
	NodeExists  map[*Node]bool
}

// CreateGraph - creates the initial graph with the truth node as
// the starting node
func CreateGraph() *Graph {

	g := Graph{
		Edges:       make(map[*Node][]*Node),
		TreeNodeMap: make(map[*TreeNode]*Node),
		NodeExists:  make(map[*Node]bool)}

	truthNode := Node{Truth: true}
	g.AddNode(&truthNode)
	return &g
}

// AddNode add a node to the graph
func (g *Graph) AddNode(n *Node) {

	if !g.NodeExists[n] {
		if n.Truth {
			g.TruthNode = n
		}

		g.Nodes = append(g.Nodes, n)
		g.NodeExists[n] = true
	}

}

func (g *Graph) AddEdge(from *Node, to *Node, weight int) {
	to.PathCost = from.PathCost + weight
	g.Edges[from] = append(g.Edges[from], to)
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

func (g *Graph) MapTreeNode(t *TreeNode) *Node {
	return g.TreeNodeMap[t]
}
