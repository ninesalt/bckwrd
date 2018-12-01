package main

import "fmt"

func main() {

	x := Node{3}
	y := Node{20}
	z := Node{15}

	g := Graph{}
	g.AddNode(&x, nil, 0)
	g.AddNode(&y, &x, 50)
	g.AddNode(&z, &y, 70)

	fmt.Println(g.Nodes[0]) //root
	fmt.Println(g.Nodes[1])
	fmt.Println(g.Nodes[2])
}
