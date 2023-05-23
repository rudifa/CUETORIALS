# Cuetorials code samples

[Cuetorials > The Go API > The Basics](https://cuetorials.com/go-api/basics/)

A compilation by [@rudifa](https://github.com/rudifa).

This directory contains sample code from `Cuetorials > The Go API > The Basics`.

There is one go project for each section of the **Basics**.

The separate `main.go` examples found in each section have been folded into a single `main.go`, with each example moved into its own function.

Imported cue packages:


```
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/format"
```
The family tree:

```
BASICS % tree                                                                           [main L|…7]
.
├── building-values
│   ├── go.mod
│   └── main.go
├── cue-context
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── error-handling
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── evaluating-values
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go-codec
│   ├── go.mod
│   └── main.go
├── inspecting-values
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── value.cue
├── printing-with-go-and-cue-format
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── traversing-values
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── value.cue
└── README.md

9 directories, 25 files

```




