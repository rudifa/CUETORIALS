package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
)

func main() {
	fmt.Println("Here we go")
	ValueValidate()
	ValueUnify()
	ValueSubsume()
	ValueEval()
}

// Value.Validate

const val = `
i: int
s: string
t: [string]: string
_h: int
_h: "hidden"
#d: int
#d: "bar"
`

func ValueValidate() {
	fmt.Println("--- Value.Validate:")
	c := cuecontext.New()
	v := c.CompileString(val)

	//   try out different validation schemes
	printErr("loose error", loose(v))
	printErr("every error", every(v))
	printErr("strict error", strict(v))

	fmt.Printf("\nvalue:\n%#v\n", v)
}

func loose(v cue.Value) error {
	return v.Validate(
		// not final or concrete
		cue.Concrete(false),
		// check minimally
		cue.Definitions(false),
		cue.Hidden(false),
		cue.Optional(false),
	)
}

func every(v cue.Value) error {
	return v.Validate(
		// not final or concrete
		cue.Concrete(false),
		// check everything
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	)
}

func strict(v cue.Value) error {
	return v.Validate(
		// ensure final and concrete
		cue.Final(),
		cue.Concrete(true),
		// check everything
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	)
}

const schema = `
v: {
	i: int
	s: string
}
`

const val2 = `
v: {
	i: "hello"
	s: 1
}
`

func ValueUnify() {
	fmt.Println("--- Value.Unify:")
	c := cuecontext.New()
	s := c.CompileString(schema, cue.Filename("schema.cue"))
	v := c.CompileString(val2, cue.Filename("val.cue"))

	// unify the schema and value
	u := s.Unify(v)

	// check for errors during unification
	if u.Err() != nil {
		msg := errors.Details(u.Err(), nil)
		fmt.Printf("Unify Error:\n%s\n", msg)
	}

	// To get all errors, we need to validate
	err := u.Validate()
	if err != nil {
		msg := errors.Details(err, nil)
		fmt.Printf("Validate Error:\n%s\n", msg)
	}

	// print u
	fmt.Printf("%#v\n", u)
}

// Value.Subsume

const schemaWithNumber = `
{
	i: number
	s: string
}
`

const schemaWithInt = `
{
	i: int
	s: string
}
`

const constraint = `
{
	i: >10 // this will only be subsumed by number, not int
	s: =~"^foo"
}
`

const val3 = `
{
	i: 100
	s: "foobar"
}
`

const constraintType = `
{
	i: uint
	s: string
}
`

func ValueSubsume() {
	fmt.Println("--- ValueSubsume:")
	ctx := cuecontext.New()
	sn := ctx.CompileString(schemaWithNumber, cue.Filename("schema_number.cue"))
	si := ctx.CompileString(schemaWithInt, cue.Filename("schema_int.cue"))
	c := ctx.CompileString(constraint, cue.Filename("constraint.cue"))
	v := ctx.CompileString(val3, cue.Filename("val.cue"))
	b := ctx.CompileString(constraintType, cue.Filename("bad.cue"))

	// check subsumptions
	printErr("sn > c", sn.Subsume(c))
	printErr("c > sn", c.Subsume(sn))

	printErr("sn > v", sn.Subsume(v))
	printErr("v > sn", v.Subsume(sn))

	// this seems not intuitive, we'll talk it later
	printErr("si > c", si.Subsume(c))
	printErr("c > si", c.Subsume(si))

	printErr("si > v", si.Subsume(v))
	printErr("v > si", v.Subsume(si))

	printErr("s > b", si.Subsume(b))
	printErr("b > v", b.Subsume(v))
}

// Value.Eval

const val4 = `
v: {
	i: int
	s: "hello"
	#d: "defn"
	_h: "hidden"
	o?: string
}
`

func ValueEval() {
	fmt.Println("--- ValueEval:")
	c := cuecontext.New()
	v := c.CompileString(val4, cue.Filename("val.cue"))

	// check for errors during compiling
	if v.Err() != nil {
		msg := errors.Details(v.Err(), nil)
		fmt.Printf("Compile Error:\n%s\n", msg)
	}
	// print v
	fmt.Printf("%#v\n", v)

	// eval evaluates and returns a new value
	e := v.Eval()
	if e.Err() != nil {
		msg := errors.Details(e.Err(), nil)
		fmt.Printf("Eval Error:\n%s\n", msg)
	}

	// print e
	fmt.Printf("%#v\n", e)
}

// Value.Evaluate
// This function has been proposed but does not exist yet.
// It would take ...Option like Validate and Syntax and
// return a new Value after processing.

// func (v Value) Evaluate(opts ...Option) Value {...}

// If you find the new function helpful, please let the devs know
// by contributing to issue 1327.//

func printErr(prefix string, err error) {
	if err != nil {
		msg := errors.Details(err, nil)
		fmt.Printf("%s:\n%s\n", prefix, msg)
	}
}
