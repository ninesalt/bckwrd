package main

// KB - basic structure that holds the graph and
// and all the clauses
type KB struct {
	Rules      []*Formula
	Facts      []string
	Graph      *Graph
	FactExists map[string]bool
}

func CreateKB() *KB {
	return &KB{FactExists: make(map[string]bool), Graph: CreateGraph()}
}

// Ask - queries the knowledge base for something
func (kb *KB) Ask(t *Term) bool {

	//TODO
	// g := kb.Graph

	return false
}

// Tell - give the KB formulas to save
func (kb *KB) TellRule(f *Formula) {

	g := kb.Graph

	n := f.Node
	antecedent := n.Left
	consequent := n.Right

	antecNode := g.MapTreeNode(antecedent)
	consNode := g.MapTreeNode(consequent)

	if antecNode == nil {
		antecNode = &Node{Clause: antecedent}
		g.TreeNodeMap[antecedent] = antecNode
	}

	if consNode == nil {
		consNode = &Node{Clause: consequent}
		g.TreeNodeMap[consequent] = consNode

		g.AddNode(consNode)
		g.AddEdge(g.GetTruthNode(), consNode, 1)
	}

	g.AddEdge(consNode, antecNode, 1)
}

func (kb *KB) TellFact(s string) {

	if !kb.FactExists[s] {
		kb.Facts = append(kb.Facts, s)
		kb.FactExists[s] = true
	}

}
