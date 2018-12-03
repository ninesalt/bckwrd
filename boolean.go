package main

import (
	"errors"
	"strings"
)

// Numerical mapping for operations
const (
	OR      int = 0
	AND     int = 1
	IMPLIES int = 2
	IFF     int = 3
)

// string representations of operations
const (
	ORSYM  string = " \u2228 "
	ANDSYM string = " \u2227 "
	IMPSYM string = " \u21D2 "
	IFFSYM string = " \u21D4 "
	NEGSYM string = "\u00AC"
)

// TreeNode - the structure that holds boolean
//only leaves will have the Name value defined
// because all other nodes will be operator nodes
type TreeNode struct {
	Operator int
	Left     *TreeNode
	Right    *TreeNode
	Name     string
	Negated  bool
}

//Formula - structure for formulas
// where Node is the root
type Formula struct {
	Name         string
	Node         *TreeNode
	TruthMapping map[string]bool
}

//CreateFormula - Creates a formula given the variable name
func CreateFormula(name string) *Formula {
	return &Formula{Name: strings.ToUpper(name), Node: &TreeNode{Name: name}}
}

// Or - Logical Or of two formulas
func (b *Formula) Or(f *Formula) {
	b.Node = &TreeNode{Operator: OR, Left: b.Node, Right: f.Node}
}

// And - Logical and of two formulas
func (b *Formula) And(f *Formula) {
	b.Node = &TreeNode{Operator: AND, Left: b.Node, Right: f.Node}
}

//Implies - Material implication of two formulas
func (b *Formula) Implies(f *Formula) {
	b.Node = &TreeNode{Operator: IMPLIES, Left: b.Node, Right: f.Node}
}

//Iff - Material equivalence of two formulas
func (b *Formula) Iff(f *Formula) {
	b.Node = &TreeNode{Operator: IFF, Left: b.Node, Right: f.Node}
}

// Negate - Negates a formula
func (b *Formula) Negate() {
	b.Node.Negated = !b.Node.Negated
}

// SetTruthValues - Sets the truth values for all values in a formula
func (b *Formula) SetTruthValues(m map[string]bool) {
	b.TruthMapping = m
}

// Evaluate - Evaluates the formula if the truth values are defined
func (b *Formula) Evaluate() (bool, error) {

	if b.TruthMapping == nil {
		return false, errors.New("No truth mapping defined")
	}

	n := b.Node
	return b.EvaluateHelper(n), nil
}

// EvaluateHelper - Recursively evaluate the formula by
// traversing the parse tree.
// Here we XOR both the nodes and the leaves with their negation
// because NOT(A AND B) is not the same as (NOT A) AND B
func (b *Formula) EvaluateHelper(n *TreeNode) bool {

	if n.Left != nil && n.Right != nil {

		switch n.Operator {
		case OR:
			return (b.EvaluateHelper(n.Left) || b.EvaluateHelper(n.Right)) != n.Negated
		case AND:
			return (b.EvaluateHelper(n.Left) && b.EvaluateHelper(n.Right)) != n.Negated
		case IMPLIES:
			return (!b.EvaluateHelper(n.Left) || b.EvaluateHelper(n.Right)) != n.Negated
		case IFF:
			return (b.EvaluateHelper(n.Left) == b.EvaluateHelper(n.Right)) != n.Negated
		}
	}

	return b.TruthMapping[n.Name] != n.Negated //xor to negate leaf if needed
}

//ToString - Returns the formula in string form
func (b *Formula) ToString() string {
	return b.ToStringHelper(b.Node)
}

//ToStringHelper - Recursively traverse the parse tree to generate
// the formula string
func (b *Formula) ToStringHelper(n *TreeNode) string {

	if n.Left != nil && n.Right != nil {

		f := ""
		// if node should be negated
		if n.Negated {
			f += NEGSYM
		}

		switch n.Operator {

		case OR:
			return f + "(" + b.ToStringHelper(n.Left) + ORSYM + b.ToStringHelper(n.Right) + ")"
		case AND:
			return f + "(" + b.ToStringHelper(n.Left) + ANDSYM + b.ToStringHelper(n.Right) + ")"
		case IMPLIES:
			return f + "(" + b.ToStringHelper(n.Left) + IMPSYM + b.ToStringHelper(n.Right) + ")"
		case IFF:
			return f + "(" + b.ToStringHelper(n.Left) + IFFSYM + b.ToStringHelper(n.Right) + ")"
		}

	}
	//if leaf is negated
	if n.Negated {
		return "(" + NEGSYM + n.Name + ")"
	}
	return n.Name
}

//Reduce - Simplifies a formula to an equivalent formula
func (b *Formula) Reduce(f *Formula) {
	//TODO
}
