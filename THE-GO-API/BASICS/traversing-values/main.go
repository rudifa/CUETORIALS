package main

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	fmt.Println("Here we go")
	Selectors_MakePath()
	ListIteration()
	FieldIteration()
	TypeSwitching()
	WalkingValue()
	CustomWalk()
}

// 	Selectors and MakePath
// We saw how to use LookupPath with ParsePath in a previous section.
// We can programmatically construct paths with Selectors and MakePath.
// We’ll also use this to reconstruct the label for the current value.

func Selectors_MakePath() {

	fmt.Println("--- Selectors_MakePath:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	// printing the label vs path
	sub := val.LookupPath(cue.ParsePath("obj.sub"))
	fmt.Println(sub.Path(), getLabel(sub))
}

// helper function for getting the label for a value
func getLabel(val cue.Value) cue.Selector {
	ss := val.Path().Selectors()
	s := ss[len(ss)-1]
	return s
}

// List Iteration

func ListIteration() {

	fmt.Println("--- ListIteration:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	// lookup a list value
	val = val.LookupPath(cue.ParsePath("obj.list"))

	// we use iterators to traverse a list
	// List() returns an iterator
	iter, _ := val.List()

	// This pattern is standard iteration
	// We get the current element and nil at the end
	for iter.Next() {
		fmt.Println(iter.Value())
	}
}

func FieldIteration() {

	fmt.Println("--- FieldIteration:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)
	val = val.LookupPath(cue.ParsePath("obj"))

	// without options
	fmt.Println("Without\n---------")
	printFields(val.Fields())

	// default options
	fmt.Println("Default\n---------")
	printFields(val.Fields(defaultOptions...))

	// custom options
	fmt.Println("Custom\n---------")
	printFields(val.Fields(customOptions...))

}

func printFields(iter *cue.Iterator, err error) {
	for iter.Next() {
		fmt.Printf("%v: %v\n", iter.Selector(), iter.Value())
	}
	fmt.Println()
}

// Type Switching
// You will likely want to make choices based on the type of a value.
// Use a switch statement on val.IncompleteKind().

func TypeSwitching() {

	fmt.Println("--- TypeSwitching:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)
	val = val.LookupPath(cue.ParsePath("obj"))

	fmt.Println("obj:")
	iter, _ := val.Fields()
	for iter.Next() {
		printNodeType(iter.Value())
	}

	fmt.Println("\nsub:")
	val = val.LookupPath(cue.ParsePath("sub"))
	iter, _ = val.Fields()
	for iter.Next() {
		printNodeType(iter.Value())
	}

}

func printNodeType(val cue.Value) {
	switch val.IncompleteKind() {
	case cue.StructKind:
		fmt.Println("struct")

	case cue.ListKind:
		fmt.Println("list")

	default:
		printLeafType(val)
	}
}

func printLeafType(val cue.Value) {
	fmt.Println(val.IncompleteKind())
}

// Walking a Value

func WalkingValue() {

	fmt.Println("--- WalkingValue:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	// before (pre-order) traversal
	preprinter := func(v cue.Value) bool {
		fmt.Printf("%v\n", v)
		return true
	}

	// after (post-order) traversal
	cnt := 0
	postcounter := func(v cue.Value) {
		cnt++
	}

	// walk the value
	val.Walk(preprinter, postcounter)

	// print count
	fmt.Println("\n\nCount:", cnt)
}

// Custom Walk
// In the previous example for default walk, some of the fields were not traversed.
// This is because CUE’s default Walk() uses the same default Field() options on a value.
// In order to walk all fields, we need to write a custom walk function where we can pass in the options for Field().

const input = `
a: {
	i: int
	j: int | *i
}
`

func CustomWalk() {

	fmt.Println("--- CustomWalk:")

	c := cuecontext.New()
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	// before (pre-order) traversal
	preprinter := func(v cue.Value) bool {
		fmt.Printf("%v\n", v)
		return true
	}

	// after (post-order) traversal
	cnt := 0
	postcounter := func(v cue.Value) {
		cnt++
	}

	Walk(val, preprinter, postcounter, customOptions...)

}

// Walk is an alternative to cue.Value.Walk which handles more field types
// You can customize this with your own options
func Walk(v cue.Value, before func(cue.Value) bool, after func(cue.Value), options ...cue.Option) {

	// call before and possibly stop recursion
	if before != nil && !before(v) {
		return
	}

	// possibly recurse
	switch v.IncompleteKind() {
	case cue.StructKind:
		if options == nil {
			options = defaultOptions
		}
		s, _ := v.Fields(options...)

		for s.Next() {
			Walk(s.Value(), before, after, options...)
		}

	case cue.ListKind:
		l, _ := v.List()
		for l.Next() {
			Walk(l.Value(), before, after, options...)
		}

		// no default (basic lit types)

	}

	if after != nil {
		after(v)
	}

}

// Cue's default
var defaultOptions = []cue.Option{
	cue.Attributes(true),
	cue.Concrete(false),
	cue.Definitions(false),
	cue.DisallowCycles(false),
	cue.Docs(false),
	cue.Hidden(false),
	cue.Optional(false),
	cue.ResolveReferences(false),
	// The following are not set
	// nor do they have a bool arg
	// cue.Final(),
	// cue.Raw(),
	// cue.Schema(),
}

// Our custom options
var customOptions = []cue.Option{
	cue.Definitions(true),
	cue.Hidden(true),
	cue.Optional(true),
}
