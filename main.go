package main

import (
	"fmt"
)

func main() {

	a := CreateFormula("A")
	b := CreateFormula("B")
	c := CreateFormula("C")
	d := CreateFormula("D")

	f := a.And(b).Or(c).Implies(d).And(d).Negate()

	mapping := make(map[string]bool)
	mapping["A"] = true
	mapping["B"] = true
	mapping["C"] = false
	mapping["D"] = false

	f.SetTruthValues(mapping)
	result, _ := f.Evaluate()
	fmt.Println(f.ToString())
	fmt.Println()
	fmt.Printf("Formula is %v \n", result)

	frog := CreateFormula("frog(X)")
	green := CreateFormula("green(X)")

	rule1 := frog.Implies(green)
	rule2 := CreateFormula("croaks(X)").Implies(frog)

	fmt.Println(rule1.ToString())
	fmt.Println(rule2.ToString())

	kb := CreateKB()

	kb.TellRule(rule1) //frog(X) -> green(X)
	kb.TellRule(rule2) //croaks(X) -> frog(X)

	kb.TellFact("croaks(fritz)")
	// t := CreateTerm("green(fritz)")
	// r := kb.Ask(t)
	// fmt.Println(r)

	t1 := CreateTerm("hello(X, f(G), Z)")
	t2 := CreateTerm("hello(a, B, f(a))")
	valid, unified := t2.UnifyTerms(t1)

	if valid {
		fmt.Printf("%v and %v unify to: %v", t1.ToString(),
			t2.ToString(), unified.ToString())
	} else {
		fmt.Println("cannot unify terms")
	}
}
