package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	fmt.Println("Here we go")
	compileWithContext()
	compileWithScope()
	encodingValuesFromGo()
}

const val = `
i: int
s: "hello"
`

// Compiling with Context

func compileWithContext() {
	var (
		c *cue.Context
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.CompileString(val)

	fmt.Println("--- compileWithContext:")

	// print the value
	fmt.Println(v)
}
func compilingWith() {
	var (
		c *cue.Context
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.CompileString(val)

	// print the value
	fmt.Println(v)
}

// Compiling with a Scope

func compileWithScope() {
	// If you have schemas in one string and values in another,
	// you can use the cue.Scope option to provide a “context” for the cue.Context.

	const schema = `
#schema: {
	i: int
	s: string
}
`

	const val2 = `
v: #schema & {
	i: 1
	s: "hello"
}
`

	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile our schema first
	s = c.CompileString(schema)

	// compile our value with scope
	v = c.CompileString(val2, cue.Scope(s))

	// print the value
	fmt.Println("--- compileWithScope:")
	fmt.Println(v)
}

type Val struct {
	I int    `json:"i"`
	S string `json:"s,omitempty"`
	b bool
}

func encodingValuesFromGo() {
	var (
		c *cue.Context
		v cue.Value
	)

	val := Val{
		I: 1,
		S: "hello",
		b: true,
	}

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.Encode(val)

	fmt.Println("--- encodingValuesFromGo:")

	// print the value
	fmt.Println(v)

	// we can also encode types
	t := c.EncodeType(Val{})

	fmt.Println(t)
}
