package main

import "fmt"

// KB - basic structure that holds the graph and
// and all the clauses
type KB struct {
	Rules []*Formula
	Facts []string
	Graph *Graph
}

// Ask - queries the knowledge base for something
// currently this only works for simple atoms (no if conditions or conjunctions, etc)
func (kb *KB) Ask(f Formula) bool {

	// g := kb.Graph
	return false
}

// Tell - give the KB a list of atoms/clauses to save
// currently this only works with simple atomic formulas
func (kb *KB) Tell(f *Formula) {

	if kb.Graph == nil {
		kb.Graph = CreateGraph()
	}

	g := kb.Graph

	// if formula is a rule (A Implies B)
	n := f.Node
	antecedent := n.Left
	consequent := n.Right

	consNode := Node{Clause: consequent}
	antecNode := Node{Clause: antecedent}

	// if consNode exists
	if g.NodeExists(&consNode) {
		fmt.Println("already exists")
	}

	g.AddNode(&consNode, g.GetTruthNode(), 1)
	g.AddNode(&antecNode, &consNode, 1)
	// fmt.Println(g)

}
