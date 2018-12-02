package main

import "fmt"

func main() {

	// x := Node{}
	// y := Node{}
	// z := Node{}

	// g := Graph{Edges: make(map[*Node][]*Node)}
	// g.AddNode(&x, nil, 0)
	// g.AddNode(&y, &x, 50)
	// g.AddNode(&z, &y, 70)

	// fmt.Println(g.Nodes[0]) //root
	// fmt.Println(g.Nodes[1])
	// fmt.Println(g.Nodes[2])

	kb := KB{}

	x := Atom{"animal", []string{"cat"}}
	y := Atom{"man", []string{"mortal"}}
	z := Atom{"animal", []string{"dog"}}

	kb.Tell([]Atom{x, y, z})

	test := Atom{"animal", []string{"dog"}}
	test2 := Atom{"animal", []string{"cow"}}

	fmt.Println(kb.Ask(test))  //true
	fmt.Println(kb.Ask(test2)) //false

	a := CreateFormula("A")
	b := CreateFormula("B")
	c := CreateFormula("C")
	d := CreateFormula("D")

	a.And(b)
	a.Or(c)
	a.Implies(d)
	// ((a AND B) OR C) IMPLIES D

	mapping := make(map[string]bool)
	mapping["A"] = true
	mapping["B"] = false
	mapping["C"] = true
	mapping["D"] = false

	a.SetTruthValues(mapping)
	result, _ := a.Evaluate()
	fmt.Printf("Formula is %v \n", result)

}
