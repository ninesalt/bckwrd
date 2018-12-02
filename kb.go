package main

import (
	"fmt"
	"reflect"
)

// KB - basic structure that holds the graph and
// and all the clauses
type KB struct {
	Clauses []*Clause
	Graph   *Graph
}

// Ask - queries the knowledge base for something
// currently this only works for simple atoms (no if conditions or conjunctions, etc)
func (kb *KB) Ask(a Atom) bool {

	g := kb.Graph
	t := g.GetTruthNode()
	reachable := g.GetConnectedNodes(t)

	for _, node := range reachable {
		if reflect.DeepEqual(node.Clause, a) {
			return !node.Truth
		}
	}
	return false
}

// Tell - give the KB a list of atoms/clauses to save
// currently this only works with simple atomic formulas
func (kb *KB) Tell(a []Atom) {

	if kb.Graph == nil {
		kb.Graph = CreateGraph()
	}

	g := kb.Graph

	for _, atom := range a {

		if !g.TypeExists[atom.Name] {
			fmt.Println("node doesnt exist")
			n := Node{Clause: atom}
			g.AddNode(&n, g.GetTruthNode(), 1)

		} else {
			fmt.Println("Node exists!")
		}
	}
}
