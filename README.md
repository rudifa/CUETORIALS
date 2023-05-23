# Cuetorials

Sample go programs from [Cuetorials](https://cuetorials.com/)

A compilation by [@rudifa](https://github.com/rudifa).

There is one go project here for each Cuetorials section where go projects were found.

Where a Cuetorials section contains several projects, their code was folded into corresponding functions, each called from the `main.go`.

Imports seen in the tutorial code:

```
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/tools/flow"
	"cuelang.org/go/cue/format"
```

Overview:

```
CUETORIALS % tree                                                                                                                            [main L|…1]
.
├── THE-GO-API
│   ├── BASICS
│   │   ├── building-values
│   │   │   ├── go.mod
│   │   │   └── main.go
│   │   ├── cue-context
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   └── main.go
│   │   ├── error-handling
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   └── main.go
│   │   ├── evaluating-values
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   └── main.go
│   │   ├── go-codec
│   │   │   ├── go.mod
│   │   │   └── main.go
│   │   ├── inspecting-values
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── main.go
│   │   │   └── value.cue
│   │   ├── printing-with-go-and-cue-format
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   └── main.go
│   │   ├── traversing-values
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── main.go
│   │   │   └── value.cue
│   │   ├── README.md
│   │   └── nitpick.txt
│   ├── CUSTOM-ATTRIBUTES
│   │   └── README.md
│   ├── LOADING-CUE
│   │   ├── loading-and-printing-cue-code
│   │   │   ├── go.mod
│   │   │   ├── go.sum
│   │   │   ├── hello.cue
│   │   │   └── main.go
│   │   └── README.md
│   ├── REPLICATING-THE-CLI
│   │   └── README.md
│   └── THE-FLOW-ENGINE
│       ├── custom-workflow-tasks
│       │   ├── go.mod
│       │   ├── go.sum
│       │   └── main.go
│       └── README.md
└── README.md

```




