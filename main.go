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

	a := Atom{"animal", []string{"cat"}}
	b := Atom{"man", []string{"mortal"}}
	c := Atom{"animal", []string{"dog"}}

	kb.Tell([]Atom{a, b, c})

	test := Atom{"animal", []string{"dog"}}
	test2 := Atom{"animal", []string{"cow"}}

	fmt.Println(kb.Ask(test))  //true
	fmt.Println(kb.Ask(test2)) //false

}
