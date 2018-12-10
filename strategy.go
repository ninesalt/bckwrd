package main

import "sort"

// GetPathCost - f(n)
func (g *Graph) GetPathCost(n *Node) int {
	return n.PathCost
}

// GetHeuristicValue - g(n)
func (g *Graph) GetHeuristicValue(n *Node) int {
	return len(g.Edges[n]) //the number of nodes this node goes to
}

//AS - A* graph search algorithm
func (g *Graph) AS(NodeQueue []*Node, ExpandedNodes []*Node) {

	NodeQueue = append(NodeQueue, ExpandedNodes...)

	// sort nodes by f(n) + g(n)
	sort.Slice(NodeQueue, func(i, j int) bool {
		n1 := g.GetPathCost(NodeQueue[i]) + g.GetHeuristicValue(NodeQueue[i])
		n2 := g.GetPathCost(NodeQueue[j]) + g.GetHeuristicValue(NodeQueue[j])
		return n1 < n2
	})

}
