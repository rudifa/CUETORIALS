# Discovering Cue and Cue Go API

My task: create a command-line app that relies on the Cue Go API to implement a subset of Cue features and use them for whatever the app has to do (something like creating end-user apps from a set of templates and configuration files and data pulled from specific remote databases).

It a appears that Cue offers most of the capabilities that our app needs, to  
validate the incoming data, to refine, process and transform them along with the constraints built into the app.

Ideally we want a single monolitic app, built for the alternate platforms, that we can deliver to end users without asking them to install or installing for them a bunch of additional tools (such as the cue cli).

So, what are our options?

- bash script invoking the cue cli and other tools
- a go program executing calls to the cue cli, e.t.c.
- a go program calling the go cli API

The first option can be of use during the early phases of the app development.

The second options is closer to what we want, but still unsuitable for the end users.

The third option is clearly what we want. It is also the most work intensive, since we need to master not only the Cue features and their usage, but also the intricacies of the Cue Go API.

So, off to a deep dive into the Cue docs and Cuetorials.

Off the bat, Cuetorials teach us how to load a .cue file into a small go app and print it out, after validation. But, how to load two .cue files, or a .cue and a .json and unify them, we have yet to find out.

Based on the Cuetorials, it appears that the Cue Go API is not a clearly outlined set of functions with clearly documented responsibilities, that would implement the equivalent of the cli commands (def, eval, export, vet, ...).

One would imagine that the cue cli commands are implemented so that they collect the user inputs (files, options and their arguments), pack them into a Request and fire off a call to the corresponding worker function, something like CueDef(r Request), CueEval(r Request), e.t.c

_Actually, the buildPlan seems to have the role of my Request._

_// A buildPlan defines what should be done based on command line
// arguments and flags._


However, the study of the source code shows that the cue commands code does all the input collection work, but it also contains the logic needed to decide which lower-level functions to call. Moreover, it uses several interesting looking structs like ...

Clone and reshape the API to suit our purposes, without breaking anything? This looks like a huge amount of work. Besides, we probably need only a subset of the cue cli functionality.

So, let's go back to the Cuetorials and Cue docs, and try to understand the mechanics and the spirit of the API.

Cue Value emerges quickly as the central entity to which the inputs and options converge and the outputs flow. However, the paths leading there are complex.

A first experiment with loading a .cue and a .json ... did not pan out.

So, am I on a wrong track? Should I study and use the cue import directive inside a cue file to import the json data? ... not clear.

Another look at the Cuetorials examples, working on strings defined within the code shows a way. So, if I read each file and Unify the resulting Values, I might be on a feasible track.

...

- I would do well if I listed all functions used in the Cuetorials examples and study the docs for each one, then apply them.

like, what's `cue.Filename("val.cue")`

Here is a start, in the order of appearance in my attention stream:

| function call| args | result | comment |
| :------- | :------- | :------- | :------- |
| c := cuecontext.New() | - | cue.Context | create a context |
| v := c.CompileString(schema) | string | cue.Value | compile a string |
| v := c.CompileString(val2, cue.Scope(v1)) | string, BuildOption | cue.Value | compile a string with scope |
| bop := cue.Scope(v) | cue.Value | cue.BuildOption | A BuildOption defines options for the various build-related methods of Context.|
| cue.Filename("sch.cue") | string | cue.BuildOption | Filename assigns a filename to parsed content. |
| ParsePath | - | - | - |
| Unify | - | - | - |
| Subsume | - | - | - |
| format.Node(val.Syntax()) | - | - | - |
| CompileBytes | - | - | - |
| - | - | - | - |
| - | - | - | - |
| - | - | - | - |
| - | - | - | - |
| - | - | - | - |
| - | - | - | - |
| - | - | - | - |

Above is a small sampling of functions seen, to say nothing of structs that are carried around - context, options, buildPlan, cue Value...

It is difficult to find in the cue code places where decisions are made and results produced, or not.


## Cue command options

| flag | info | def | eval | export | import | vet |
| :------- | :------- | :------- | :------- | :------- | :------- |:------- |
| --all| show optional and hidden fields|- | y| -| -|- |
| --concrete| require the evaluation to be concrete| -|y | -| -| y|
| --dryrun | only run simulation| -|- |- | y|- |
| --escape | use HTML escaping| -|- |y | -| -|
| --expression | evaluate this expression only | y | y| y| -| -|
| --ext| match files with these extensions| -| -|- |y | -|
| --force | force overwriting existing files |  y | y| -| y|- |
| --help | help for def | y |y|y| y| y|
| --inject  | set the value of a tagged field | y | y|y |- |y |
| --inject-vars  | inject system variables in tags | y |y | y| -| y|
| --inline-imports | expand references to non-core imports | y |- | -| -| -|
| --list | concatenate multiple objects into a list | y | y| y| y| y|
| --merge   | merge non-CUE files (default true) | y | y| y| y| y|
| --name | glob filter for non-CUE file names in directories | y | y| y| y| y|
| --out | output format (run 'cue filetypes' for more info) | y | y| y| -| -|
| --outfile|filename or - for stdout with optional file prefix (run 'cue filetypes' for more info) | y| y|y | y| -|
| --package | package name for non-CUE files | y | y| y| y|y |
| --path | CUE expression for single path component | y | y| y| y| y|
| --proto_enum | mode for rendering enums (int \| json) (default "int") | y | y| y|y | y|
| --proto_path | paths in which to search for imports| y | y| y| y| y|
| --schema | expression to select schema for evaluating values in non-CUE files | y | y| y| y|y |
| --show-attributes | display field attributes | y | y|- |- | -|
| --show-hidden | display hidden fields|- | y| -|- | -|
| --show-optional |display optional fields |-|y | -| -| -| -| -|
| --with-context | import as object with contextual data | y |y | y| y| y|
| | | | | | | |


## Investigating

	
- My immediate objective: write a minimal replica of the `cue vet`command, supporting the `options --concrete` and `--schema <schema def>` and expecting a `.cue` schema file and a `.json` data file.

The first naive attempt at unifying a schema with the .cue and .json files is no good.

What next?

Stepping into cue with the debugger works, but it is difficult to see just what is happenning and where.

Read the code? The general docs?

### The tutorial samples are few, and may use deprecated funcs or types

So, what should we do?

Start collecting the approved (not deprecated) code snippets, assorted with tests, and use those for the further app development.


### The Cue Go API: what is it?

- announcements...
- indeed, third parties build on it: hofstetter, dragon, others?

So, what is the extent of the Cue Go API v0.5.0?

- the set of functions and types shown in the official tutorials (excluding or replacing the deprecated ones)?
- the set of functions and types used by the cue command implemementation?
- the set of functions and types documented for all cue packages that are not internal?


From [docs-integrations-go](https://cuelang.org/docs/integrations/go/):

The CUE APIs in the main repo are organized as follows:

- `cmd`: The CUE command line tool.
- `cue`: core APIs related to parsing, formatting, loading and running CUE programs. These packages are used by all other packages, including the command line tool.
- `encoding` Go, Protobuf, and OpenAPI.
- `pkg`: Builtin packages that are available from within CUE programs. These are typically not used in Go code.

Based on above, we should concentrate on `cue` and `encoding` packages.

### from cue docs

> Overview 
> 
>Package cue is the main API for CUE evaluation.

>Value is the main type that represents CUE evaluations. Values are created with a cue.Context. Only values created from the same Context can be involved in the same operation.

>A Context defines the set of active packages, the translations of field names to unique codes, as well as the set of builtins.

### Cue operations

- vet / validate 1.json 1.cue
- vet / validate 1.json 1.cue -d #def
- extract a subset of validated data
- transform extracted data to a digest .cue or .json or .yaml
- merge extracted data from several inputs

### Learning directions
- examples of Cue operations using the cue clî
- porting these operations to go programs
- reading cue code
- grepping cue code
- debugging cue code
- using dependency viewers
- implementing the required operations in the target app
- distilling util code into packages for reuse

### Read a .json file into a cue Value

- a go program using cue go api: what is the proper code to read a .json file into a cue Value

```
package main

import (
	"fmt"
	"io/ioutil"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
)

func main() {
	// Read the JSON file
	jsonBytes, err := ioutil.ReadFile("example.x.json")
	if err != nil {
		fmt.Printf("Error reading JSON file: %v", err)
		return
	}

	// Create a new Cue context
	ctx := cuecontext.New()

	// Decode the JSON data into a cue.Value
	val := ctx.CompileBytes(jsonBytes)

	err = val.Err()
	if err != nil {
		fmt.Printf("Error decoding JSON: %v", err)
		return
	}

	// Format the cue.Value for pretty printing
	formatted, err := format.Node(val.Syntax())
	if err != nil {
		fmt.Printf("Error formatting cue.Value: %v", err)
		return
	}

	fmt.Println(string(formatted))
}
```
Questions:

- whence comes a bad json diagnôsstic?

`cue def bad.json` flags it: `invalid JSON for file...`,?

```
cue % grep.go 'invalid JSON for file'                                      [try-split-off-api L|…1]
encoding/json/json.go:90:		return nil, errors.Wrapf(err, p, "invalid JSON for file %q", path)
encoding/json/json.go:138:		return nil, errors.Wrapf(err, pos, "invalid JSON for file %q", d.path)
encoding/json/json_test.go:107:		out:  "invalid JSON for file \"invalid JSON\": invalid character '_' after array element",
```
So, I should look to `encoding/json` package.

- whence comes a bad cue diagnostic?

```
cue % grep.go 'reference .* not found'                                                                                            [try-split-off-api L|…1]
internal/core/compile/compile.go:393:		return c.errf(n, "reference %q not found", n.Name)
```

```
cue % grep.go 'string literal not terminated'                                                                                
cue/path_test.go:102:		out:  `_|_ // string literal not terminated`,
cue/scanner/scanner.go:464:			s.errf(offs, "string literal not terminated")
cue/scanner/scanner_test.go:777:	// {`"\u000`, token.STRING, 6, `"\u000`, "string literal not terminated"}, two errors
...
```

- `ctx.CompileString(string(bytes))` vs `ctx.CompileBytes(bytes)`: what difference?


```
// CompileString parses and build a Value from the given source string.
//
// The returned Value will represent an error, accessible through Err, if any
// error occurred.
func (c *Context) CompileString(src string, options ...BuildOption) Value {
	cfg := c.parseOptions(options)
	return c.compile(c.runtime().Compile(&cfg, src))
}

// CompileBytes parses and build a Value from the given source bytes.
//
// The returned Value will represent an error, accessible through Err, if any
// error occurred.
func (c *Context) CompileBytes(b []byte, options ...BuildOption) Value {
	cfg := c.parseOptions(options)
	return c.compile(c.runtime().Compile(&cfg, b))
}
```

- `ctx.CompileBytes(jsonBytes)` seems to tolerate malformed json, so we need to validate json separately?
- `format.Node(val.Value().Syntax()` and `format.Node(val.Syntax()` seem to do  same thing ?!
- can we parametrize `format.Node` to print compact representation?


This looks like the portal to file entry

```
cue % tree pkg/encoding                                                                             [try-split-off-api L|…1]
pkg/encoding
├── base64
│   ├── testdata
│   │   └── gen.txtar
│   ├── base64_test.go
│   ├── manual.go
│   └── pkg.go
├── csv
│   ├── testdata
│   │   └── gen.txtar
│   ├── csv_test.go
│   ├── manual.go
│   └── pkg.go
├── hex
│   ├── testdata
│   │   └── gen.txtar
│   ├── hex.go
│   ├── hex_test.go
│   ├── manual.go
│   └── pkg.go
├── json
│   ├── testdata
│   │   └── gen.txtar
│   ├── json.go
│   ├── json_test.go
│   ├── manual.go
│   └── pkg.go
└── yaml
    ├── testdata
    │   └── gen.txtar
    ├── manual.go
    ├── pkg.go
    └── yaml_test.go
```
What, no package for .cue files?


