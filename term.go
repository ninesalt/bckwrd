package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Term - recursively defined data structure to hold variables and
// constants of any depth : a(b(c, g), f(h(p)), ..)
type Term struct {
	IsConstant     bool
	IsVar          bool
	IsFunc         bool
	Predicate      string
	PredicateTerms []*Term
	StringRep      string
}

// CreateTerm recursively creates the terms
func CreateTerm(s string) *Term {

	s = strings.TrimSpace(s)
	r := regexp.MustCompile("\\w+\\(.+\\,*.*\\)")
	isFunc := r.MatchString(s)
	t := Term{StringRep: s}

	if isFunc {

		PredRegex := regexp.MustCompile("\\w+")
		p := PredRegex.FindAllString(s, 1)[0]
		t.Predicate = p

		firstIndex := strings.Index(s, "(")
		lastIndex := strings.LastIndex(s, ")")
		argsString := s[firstIndex+1 : lastIndex]
		args := strings.Split(argsString, ",")
		t.IsFunc = true

		for _, a := range args {
			t.PredicateTerms = append(t.PredicateTerms, CreateTerm(a))
		}

	} else {
		_, err := strconv.Atoi(s)
		IsNum := err == nil
		first := string(s[0])
		IsVar := !IsNum && strings.ToUpper(first) == first
		t.IsVar = IsVar
		t.Predicate = s
		t.IsConstant = !IsVar
	}

	return &t
}

//ToString - returns the string representation of a term
func (t *Term) ToString() string {

	if t.StringRep != "" {
		return t.StringRep
	}

	if !t.IsVar && !t.IsConstant {
		p := t.Predicate
		var Preds []string

		for _, g := range t.PredicateTerms {
			Preds = append(Preds, g.ToString())
		}

		return fmt.Sprintf("%v(%v)", p, strings.Join(Preds, ", "))
	}

	return t.Predicate

}

// UnifyTerms - basic unification function that unifies two
// terms together depending on their type.
// This function currently doesn't save which variable maps to
// what, as a result, it can map X to both a and b
func (t *Term) UnifyTerms(t2 *Term) (bool, *Term) {

	// cant unify 5 with 4
	if t.IsConstant && t2.IsConstant && t.Predicate != t2.Predicate {
		return false, nil
	}

	//unify Y with f(x)
	if t.IsFunc && t2.IsVar {
		ret := *t2
		return true, &ret
	}

	if t2.IsFunc && t.IsVar {
		ret := *t
		return true, &ret
	}

	// unify X with 4
	if t.IsVar && !t2.IsVar {
		return true, CreateTerm(t2.Predicate)
	}

	if !t.IsVar && t2.IsVar {
		return true, CreateTerm(t.Predicate)
	}

	// unify X with Y
	if t.IsVar && t2.IsVar || t.ToString() == t2.ToString() {
		ret := *t
		return true, &ret
	}

	//unify f(X, a(b)) with f(a, a(X))
	if t.Predicate == t2.Predicate &&
		len(t.PredicateTerms) == len(t2.PredicateTerms) {

		var ArgsUnified []*Term

		for i, Pred := range t.PredicateTerms {
			valid, res := Pred.UnifyTerms(t2.PredicateTerms[i])

			if !valid {
				return false, nil
			}

			ArgsUnified = append(ArgsUnified, res)
		}

		unified := Term{Predicate: t.Predicate, PredicateTerms: ArgsUnified}
		return true, &unified
	}

	return false, nil
}
