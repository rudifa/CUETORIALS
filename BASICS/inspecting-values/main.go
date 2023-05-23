package main

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	fmt.Println("Here we go")
	LookupPath_Path()
	ReferencePath_Dereference()
	Exists_IsConcrete()
	Kind_IncompleteKind()
	TypeConversions()
	Len()
}

// 	LookupPath, Path
// Path return paths to the value which can be used with LookupPath to get the value.
// They can be considered inverse functions of each other.

func LookupPath_Path() {

	fmt.Println("--- LookupPath_Path:")
	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"a",
		"d.f",
		"l",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		fmt.Printf("%q\n%# v\n", p, v)
	}
}

// ReferencePath and Dereference

func ReferencePath_Dereference() {

	fmt.Println("--- ReferencePath_Dereference:")

	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"d.f",
		"r.s",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		_, r := v.ReferencePath()
		fmt.Printf("%q %q\n%# v\n", p, r, v)
	}
}

// Exists and IsConcrete
// When you lookup a value, how do you know if it was found? That is where Exists comes in.

// IsConcrete can tell you if an atom field has data or is a terminal error.
// For lists and structs, it will report true if they exist and not recurse to check subvalues.
// When disjunctions and defaults are used…

func Exists_IsConcrete() {

	fmt.Println("--- Exists_IsConcrete:")

	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"a",
		"d.g",
		"x.y",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		x := v.Exists()
		c := v.IsConcrete()
		fmt.Printf("%q %v %v\n%# v\n", p, x, c, v)
	}
}

// Kind and IncompleteKind
// Kind and IncompleteKind will tell you the underlying type of a value.
// IncompleteKind is more granular and returns type info regarless of
// how complete a value is (the names may seem a bit backwards).

func Kind_IncompleteKind() {

	fmt.Println("--- Kind_IncompleteKind:")

	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"a",
		"a.e",
		"d.g",
		"l",
		"b",
		"x.y",
	}

	for _, path := range paths {
		v := val.LookupPath(cue.ParsePath(path))
		k := v.Kind()
		i := v.IncompleteKind()
		fmt.Printf("%q %v %v %v\n", path, k, i, v)
	}
}

// Type Conversions
// Values have a number of functions for turning the abstract into the underlying type.
// You will first want to know what type of value you are dealing with before trying to convert it.

func TypeConversions() {

	fmt.Println("--- TypeConversions:")

	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	var (
		s string
		b []byte
		i int64
		f float64
	)

	// read into Go basic types
	s, _ = val.LookupPath(cue.ParsePath("a.b")).String()
	b, _ = val.LookupPath(cue.ParsePath("b")).Bytes()
	i, _ = val.LookupPath(cue.ParsePath("d.e")).Int64()
	f, _ = val.LookupPath(cue.ParsePath("d.f")).Float64()

	fmt.Println(s, b, i, f)

	// an error
	s, err := val.LookupPath(cue.ParsePath("a.e")).String()
	if err != nil {
		fmt.Println(err)
	}

}

// Len
// Len will tell you the length of a list or how many bytes are in a bytes.

func Len() {

	fmt.Println("--- Len:")

	c := cuecontext.New()

	// read and compile value
	d, _ := os.ReadFile("value.cue")
	val := c.CompileBytes(d)

	paths := []string{
		"a",
		"d.e",
		"d.g",
		"l",
		"b",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		k := v.IncompleteKind()
		l := v.Len()
		fmt.Printf("%q %v %v\n%# v\n", p, k, l, v)
	}
}
