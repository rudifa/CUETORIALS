package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueformat "cuelang.org/go/cue/format"
)

func main() {
	fmt.Println("Here we go")
	printingWithGoFmt()
	printingWithCueFmt()
}

const val2 = `
i: int
s: string
t: [string]: string
_h: "hidden"
#d: foo: "bar"
`

func printingWithGoFmt() {
	var (
		c *cue.Context
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.CompileString(val2)

	// print the value
	fmt.Println("--- printingWithGoFmt:")
	fmt.Printf("// %%v\n%v\n\n// %%# v\n%# v\n", v, v)
}

const val = `
i: int
s: string
t: [string]: string
_h: "hidden"
#d: foo: "bar"
`

func printingWithCueFmt() {

	var (
		c *cue.Context
		v cue.Value
	)

	// create a context
	c = cuecontext.New()

	// compile some CUE into a Value
	v = c.CompileString(val)

	// Generate an AST
	//   try out different options
	syn := v.Syntax(
		cue.Final(),         // close structs and lists
		cue.Concrete(false), // allow incomplete values
		cue.Definitions(false),
		cue.Hidden(true),
		cue.Optional(true),
		cue.Attributes(true),
		cue.Docs(true),
	)

	// Pretty print the AST, returns ([]byte, error)
	bs, _ := cueformat.Node(
		syn,
		// format.TabIndent(false),
		// format.UseSpaces(2),
	)

	// print to stdout
	fmt.Println("--- printingWithCueFmt:")
	fmt.Println(string(bs))
}
