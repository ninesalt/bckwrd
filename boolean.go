package main

import (
	"errors"
)

// TreeNode - the structure that holds boolean
// expressions where operator is:
// OR: 0
// AND: 1
// IMPLIES: 2
// IFF: 3
type TreeNode struct {
	Operator   int
	TruthValue bool
	Left       *TreeNode
	Right      *TreeNode
	Name       string
}

type Formula struct {
	Name         string
	Node         *TreeNode
	TruthMapping map[string]bool
}

func CreateFormula(name string) *Formula {
	return &Formula{Name: name, Node: &TreeNode{Name: name}}
}

func (b *Formula) Or(f *Formula) {
	b.Node = &TreeNode{Operator: 0, Left: b.Node, Right: f.Node}
}

func (b *Formula) And(f *Formula) {
	b.Node = &TreeNode{Operator: 1, Left: b.Node, Right: f.Node}
}

func (b *Formula) Implies(f *Formula) {
	b.Node = &TreeNode{Operator: 2, Left: b.Node, Right: f.Node}
}

func (b *Formula) Iff(f *Formula) {
	b.Node = &TreeNode{Operator: 3, Left: b.Node, Right: f.Node}
}

func (b *Formula) Negate() {
	//TODO
}

func (b *Formula) SetTruthValues(m map[string]bool) {
	b.TruthMapping = m
}

func (b *Formula) Evaluate() (bool, error) {

	if b.TruthMapping == nil {
		return false, errors.New("No truth mapping defined")
	}

	n := b.Node
	return b.EvaluateHelper(n), nil
}

func (b *Formula) EvaluateHelper(n *TreeNode) bool {

	if n.Left != nil && n.Right != nil {

		switch n.Operator {
		case 0:
			return b.EvaluateHelper(n.Left) || b.EvaluateHelper(n.Right)
		case 1:
			return b.EvaluateHelper(n.Left) && b.EvaluateHelper(n.Right)
		case 2:
			return !b.EvaluateHelper(n.Left) || b.EvaluateHelper(n.Right)
		case 3:
			return b.EvaluateHelper(n.Left) == b.EvaluateHelper(n.Right)
		}
	}

	name := n.Name
	return b.TruthMapping[name]
}

func (b *Formula) Reduce(f *Formula) {

}
