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

So, off we swimm, towards a deep dive into the Cue docs and the Cuetorials.

Off the bat, Cuetorials teach us how to load a .cue file into a small go app and print it out, after validation. But, how to load two .cue files, or a .cue and a .json and unify them, we'll have to find out.

Based on the Cuetorials, it appears that the Cue Go API is not a clearly outlined set of functions with clearly documented responsibilities, that would implement the equivalent of the cli commands (def, eval, export, vet, ...).

One would imagine that the cue cli commands are implemented so that they collect the user inputs (files, options and their arguments), pack them into a Request and fire off a call to the corresponding worker function, something like CueDef(r Request), CueEval(r Request), e.t.c

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

Here is a start, in no particular order:

| function call| args | result | comment |
| :------- | :------- | :------- | :------- |
| c := cuecontext.New() | - | cue.Context | create a context |
| v := c.CompileString(schema) | string | cue.Value | compile a string |
| v := c.CompileString(val2, cue.Scope(v1)) | string, BuildOption | cue.Value | compile a string with scope |
| bop := cue.Scope(v) | cue.Value | cue.BuildOption | A BuildOption defines options for the various build-related methods of Context.|
| cue.Filename("sch.cue") | string | cue.BuildOption | Filename assigns a filename to parsed content. |
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
- indeed, third partios build on it: hofstetter, dragon, others?

So, what is the extent of the Cue Go API v0.5.0?

- the set of functions and types shown in the official tutorials (excluding or replacing the deprecated ones)?
- the set of functions and types used by the cue command implmementation?
