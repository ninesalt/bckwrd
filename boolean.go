package main

import (
	"errors"
)

// Numerical mapping for operations
const (
	OR      int = 0
	AND     int = 1
	IMPLIES int = 2
	IFF     int = 3
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
// where Root is the root
type Formula struct {
	Name         string
	Root         *TreeNode
	TruthMapping map[string]bool
}

//CreateFormula - Creates a formula given the variable name
func CreateFormula(name string) *Formula {
	return &Formula{Name: name, Root: &TreeNode{Name: name}}
}

// Or - Logical Or of two formulas
func (b *Formula) Or(f *Formula) {
	b.Root = &TreeNode{Operator: OR, Left: b.Root, Right: f.Root}
}

// And - Logical AND of two formulas
func (b *Formula) And(f *Formula) {
	b.Root = &TreeNode{Operator: AND, Left: b.Root, Right: f.Root}
}

//Implies - Material implication of two formulas
func (b *Formula) Implies(f *Formula) {
	b.Root = &TreeNode{Operator: IMPLIES, Left: b.Root, Right: f.Root}
}

//Iff - Material equivalence of two formulas
func (b *Formula) Iff(f *Formula) {
	b.Root = &TreeNode{Operator: IFF, Left: b.Root, Right: f.Root}
}

// Negate - Negates a formula
func (b *Formula) Negate() {
	b.Root.Negated = !b.Root.Negated
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

	n := b.Root
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
	return b.ToStringHelper(b.Root)
}

//ToStringHelper - Recursively traverse the parse tree to generate
// the formula string
func (b *Formula) ToStringHelper(n *TreeNode) string {

	if n.Left != nil && n.Right != nil {

		f := ""
		// if node should be negated
		if n.Negated {
			f += "NOT"
		}

		switch n.Operator {

		case OR:
			return f + "(" + b.ToStringHelper(n.Left) + " OR " + b.ToStringHelper(n.Right) + ")"
		case AND:
			return f + "(" + b.ToStringHelper(n.Left) + " AND " + b.ToStringHelper(n.Right) + ")"
		case IMPLIES:
			return f + "(" + b.ToStringHelper(n.Left) + " IMPLIES " + b.ToStringHelper(n.Right) + ")"
		case IFF:
			return f + "(" + b.ToStringHelper(n.Left) + " IFF " + b.ToStringHelper(n.Right) + ")"
		}

	}
	//if leaf is negated
	if n.Negated {
		return "(NOT " + n.Name + ")"
	}
	return n.Name
}

//Reduce - Simplifies a formula to an equivalent formula
func (b *Formula) Reduce(f *Formula) {
	//TODO
}
