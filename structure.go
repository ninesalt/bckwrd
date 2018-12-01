package main

// Atom - the basic building blocks of clauses
// example: animal(dog), where animal is the name
// and dog is the 0th entry in Values
type Atom struct {
	Name   string
	Values []string
}

// Clause - the conjunction/disjunction of multiple atoms
type Clause struct {
	Atoms    []*Atom
	Operator string
}
