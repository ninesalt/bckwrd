package main

import (
	"regexp"
	"strings"
)

// Term - recursively defined data structure to hold variables and
// constants of any depth : a(b(c, g), f(h(p)), ..)
type Term struct {
	IsConstant     bool
	IsVar          bool
	Predicate      string
	PredicateTerms []*Term
}

// CreateTerm recursively creates the terms
func CreateTerm(s string) *Term {

	s = strings.TrimSpace(s)
	r := regexp.MustCompile("\\w+\\(.+\\,*.*\\)")
	isFunc := r.MatchString(s)
	t := Term{}

	if isFunc {

		PredRegex := regexp.MustCompile("\\w+")
		p := PredRegex.FindAllString(s, 1)[0]
		t.Predicate = p

		firstIndex := strings.Index(s, "(")
		lastIndex := strings.LastIndex(s, ")")
		argsString := s[firstIndex+1 : lastIndex]
		args := strings.Split(argsString, ",")

		for _, a := range args {
			t.PredicateTerms = append(t.PredicateTerms, CreateTerm(a))
		}

	} else {
		first := string(s[0])
		IsVar := strings.ToUpper(first) == first
		t.IsVar = IsVar
		t.IsConstant = !IsVar
	}

	return &t
}
