package main

//Filter - filter out nodes that have already been visited
func Filter(nodes []*Node, history map[*Node]bool) {

	for i, n := range nodes {
		if history[n] {
			nodes = append(nodes[:i], nodes[i+1:]...)
		}
	}

}

func (g *Graph) GeneralSearch() bool {

	var nodes []*Node
	nodes = append(nodes, g.GetTruthNode())
	history := make(map[*Node]bool)

	for true {

		if len(nodes) == 0 {
			return false
		}
		n := nodes[0]
		nodes = nodes[1:]

		// test := GoalTest()
		// if test ...

		history[n] = true
		connections := g.GetConnectedNodes(n)
		Filter(connections, history)
		g.AS(nodes, connections)

	}
	return false
}
