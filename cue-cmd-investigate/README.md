# Investigate cue command behavior

```
cue-cmd-investigate % cat quad                                                     
#! /bin/bash
# run 4 cue commands on files and options
set -x
for i; do 
   cat $i
done
cue def $@
cue eval $@
cue export $@
cue vet $@
```

- no cue files

``` 
tmp % ../quad                                                                      
+ cue def
no CUE files in .
...
```

- empty cue file

```
cue-cmd-investigate % quad 0.cue                                                   
+ cat 0.cue
+ cue def 0.cue

+ cue eval 0.cue

+ cue export 0.cue
{}
+ cue vet 0.cue
```

- some cue files present, but not declaring a package

```
build constraints exclude all CUE files in .:
    CUETORIALS/cue-cmd-investigate/0.cue: no package name
    CUETORIALS/cue-cmd-investigate/1a.cue: no package name
...
```
- braces are implicit at the top level of a .cue file...

```
cue-cmd-investigate % quad 1*                                                    
+ cat 1a.cue
A:1
+ cat 1b.cue
B:2
+ cue def 1a.cue 1b.cue
A: 1
B: 2
+ cue eval 1a.cue 1b.cue
A: 1
B: 2
+ cue export 1a.cue 1b.cue
{
    "A": 1,
    "B": 2
}
+ cue vet 1a.cue 1b.cue
```
- ... except when the top level is a list

```
cue-cmd-investigate % quad 2a.cue                                                [load-json L|✚1…2]
+ cat 2a.cue
[1,2,3]
+ cue def 2a.cue
[1, 2, 3]
+ cue eval 2a.cue
[1, 2, 3]
+ cue export 2a.cue
[
    1,
    2,
    3
]
+ cue vet 2a.cue
```

- top level list unification example

```
cue-cmd-investigate % quad 2b.cue 2a.cue                                         [load-json L|✚2…2]
+ cat 2b.cue
[1,2,3,4]
+ cat 2a.cue
[1,2,3,...]
+ cue def 2b.cue 2a.cue

[1, 2, 3, ...]
[1, 2, 3, 4]
+ cue eval 2b.cue 2a.cue
[1, 2, 3, 4]
+ cue export 2b.cue 2a.cue
[
    1,
    2,
    3,
    4
]
+ cue vet 2b.cue 2a.cue
```

- schema example 

```
- cue-cmd-investigate % quad 100.schema.cue                                                         [load-json L|✚2…3]
+ cat 100.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/

#Conn: {
    address:  string
    port:     int
    protocol: string
}

lossy: #Conn & {
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
+ cue def 100.schema.cue
#Conn: {
	address:  string
	port:     int
	protocol: string
}
lossy: #Conn & {
	address:  "1.2.3.4"
	port:     8888
	protocol: "udp"
}
+ cue eval 100.schema.cue
#Conn: {
    address:  string
    port:     int
    protocol: string
}
lossy: {
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
+ cue export 100.schema.cue
{
    "lossy": {
        "address": "1.2.3.4",
        "port": 8888,
        "protocol": "udp"
    }
}
+ cue vet 100.schema.cue
```

- move the data to a separate file

```
cue-cmd-investigate % quad 101*                                                                     [load-json L|…4]
+ cat 101.data.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
lossy: {
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
+ cat 101.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
    address:  string
    port:     int
    protocol: string
}

lossy: #Conn 
+ cue def 101.data.cue 101.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
	address:  string
	port:     int
	protocol: string
}

// https://cuelang.org/docs/tutorials/tour/intro/schema/
lossy: #Conn & {
	address:  "1.2.3.4"
	port:     8888
	protocol: "udp"
}
+ cue eval 101.data.cue 101.schema.cue
#Conn: {
    address:  string
    port:     int
    protocol: string
}
lossy: {
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
+ cue export 101.data.cue 101.schema.cue
{
    "lossy": {
        "address": "1.2.3.4",
        "port": 8888,
        "protocol": "udp"
    }
}
+ cue vet 101.data.cue 101.schema.cue
```

- same separation, but top level data unnamed

```
cue-cmd-investigate % quad 102*                                                                   [load-json L|✚1…4]
+ cat 102.data.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
{
    address:  "1.2.3.4"
    port:     8888
    protocol: "udp"
}
+ cat 102.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
    address:  string
    port:     int
    protocol: string
}

#Conn 
+ cue def 102.data.cue 102.schema.cue

_#def
_#def: {
	#Conn
	address: "1.2.3.4"
	port:    8888
	// https://cuelang.org/docs/tutorials/tour/intro/schema/
	#Conn: {
		address:  string
		port:     int
		protocol: string
	}

	protocol: "udp"
}
+ cue eval 102.data.cue 102.schema.cue
address: "1.2.3.4"
#Conn: {
    address:  string
    port:     int
    protocol: string
}
port:     8888
protocol: "udp"
+ cue export 102.data.cue 102.schema.cue
{
    "address": "1.2.3.4",
    "port": 8888,
    "protocol": "udp"
}
+ cue vet 102.data.cue 102.schema.cue
```

- same separation, but data moved to a `.json` file

```
cue-cmd-investigate % quad 103*                                                                   [load-json L|✚6…2]
+ cat 103.data.json
{
    "address":  "1.2.3.4",
    "port":     8888,
    "protocol": "udp"
}
+ cat 103.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
    address:  string
    port:     int
    protocol: string
}

#Conn 
+ cue def 103.data.json 103.schema.cue

_#def
_#def: {
	#Conn
	address: "1.2.3.4"
	port:    8888
	// https://cuelang.org/docs/tutorials/tour/intro/schema/
	#Conn: {
		address:  string
		port:     int
		protocol: string
	}
	protocol: "udp"
}
+ cue eval 103.data.json 103.schema.cue
address: "1.2.3.4"
#Conn: {
    address:  string
    port:     int
    protocol: string
}
port:     8888
protocol: "udp"
+ cue export 103.data.json 103.schema.cue
{
    "address": "1.2.3.4",
    "port": 8888,
    "protocol": "udp"
}
+ cue vet 103.data.json 103.schema.cue
```

- same as above, except the lone `#Conn` removed 

```
cue-cmd-investigate % quad 104*                                                                     [load-json L|…2]
+ cat 104.data.json
{
    "address":  "1.2.3.4",
    "port":     8888,
    "protocol": "udp"
}
+ cat 104.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
    address:  string
    port:     int
    protocol: string
}

+ cue def 104.data.json 104.schema.cue
address: "1.2.3.4"
port:    8888
// https://cuelang.org/docs/tutorials/tour/intro/schema/
#Conn: {
	address:  string
	port:     int
	protocol: string
}
protocol: "udp"
+ cue eval 104.data.json 104.schema.cue
address: "1.2.3.4"
port:    8888
#Conn: {
    address:  string
    port:     int
    protocol: string
}
protocol: "udp"
+ cue export 104.data.json 104.schema.cue
{
    "address": "1.2.3.4",
    "port": 8888,
    "protocol": "udp"
}
+ cue vet 104.data.json 104.schema.cue
```
- make the def hidden: data are still exported, presumably because concrete and no conflict

```
cue-cmd-investigate % quad 105*                                                                     [load-json L|…2]
+ cat 105.data.json
{
    "address":  "1.2.3.4",
    "port":     8888,
    "protocol": "udp"
}
+ cat 105.schema.cue
// https://cuelang.org/docs/tutorials/tour/intro/schema/
_#Conn: {
    address:  string
    port:     int
    protocol: string
}

+ cue def 105.data.json 105.schema.cue
address: "1.2.3.4"
port:    8888
// https://cuelang.org/docs/tutorials/tour/intro/schema/
_#Conn: {
	address:  string
	port:     int
	protocol: string
}
protocol: "udp"
+ cue eval 105.data.json 105.schema.cue
address:  "1.2.3.4"
port:     8888
protocol: "udp"
+ cue export 105.data.json 105.schema.cue
{
    "address": "1.2.3.4",
    "port": 8888,
    "protocol": "udp"
}
+ cue vet 105.data.json 105.schema.cue
```
see also [--schema example](https://tidycloudaws.com/take-a-cue-to-supercharge-your-configurations/)


- schema example from [cuelang use cases](https://cuelang.org/docs/usecases/datadef/) *Validating backwards compatibility*

```
cue-cmd-investigate % quad 200.*                                                                  [load-json L|✚1…4]
+ cat 200.data.json
{
    "age": 85,
    "hobby": "software programming"
}+ cat 200.schema.cue
// https://cuelang.org/docs/usecases/datadef/
// Release notes:
// - You can now specify your age and your hobby!
#V1: {
    age:   >=0 & <=100
    hobby: string
}
// Release notes:
// - People get to be older than 100, so we relaxed it.
// - It seems not many people have a hobby, so we made it optional.
#V2: {
    age:    >=0 & <=150 // people get older now
    hobby?: string      // some people don't have a hobby
}
// Release notes:
// - Actually no one seems to have a hobby nowadays anymore, so we dropped the field.
#V3: {
    age: >=0 & <=150
}
+ cue def 200.data.json 200.schema.cue
age: 85
// https://cuelang.org/docs/usecases/datadef/
// Release notes:
// - You can now specify your age and your hobby!
#V1: {
	age:   >=0 & <=100
	hobby: string
}

// Release notes:
// - People get to be older than 100, so we relaxed it.
// - It seems not many people have a hobby, so we made it optional.
#V2: {
	age:    >=0 & <=150
	hobby?: string
}
hobby: "software programming"
// Release notes:
// - Actually no one seems to have a hobby nowadays anymore, so we dropped the field.
#V3: {
	age: >=0 & <=150
}
+ cue eval 200.data.json 200.schema.cue
age: 85
#V1: {
    age:   >=0 & <=100
    hobby: string
}
#V2: {
    age: >=0 & <=150
}
hobby: "software programming"
#V3: {
    age: >=0 & <=150
}
+ cue export 200.data.json 200.schema.cue
{
    "age": 85,
    "hobby": "software programming"
}
+ cue vet 200.data.json 200.schema.cue
```

- same schema file, select a non-conflicting schema `#V1` or `#V2`: the result is same as above

- same schema file, select a conflicting schema `#V3`

```
cue-cmd-investigate % quad 200.* -d '#V3'                                                         [load-json L|✚1…4]
+ cat 200.data.json
{
    "age": 85,
    "hobby": "software programming"
}+ cat 200.schema.cue
// https://cuelang.org/docs/usecases/datadef/
// Release notes:
// - You can now specify your age and your hobby!
#V1: {
    age:   >=0 & <=100
    hobby: string
}
// Release notes:
// - People get to be older than 100, so we relaxed it.
// - It seems not many people have a hobby, so we made it optional.
#V2: {
    age:    >=0 & <=150 // people get older now
    hobby?: string      // some people don't have a hobby
}
// Release notes:
// - Actually no one seems to have a hobby nowadays anymore, so we dropped the field.
#V3: {
    age: >=0 & <=150
}
+ cue def 200.data.json 200.schema.cue -d '#V3'
hobby: field not allowed:
    ./200.data.json:3:5
    ./200.schema.cue:17:6
+ cue eval 200.data.json 200.schema.cue -d '#V3'
hobby: field not allowed:
    ./200.data.json:3:5
    ./200.schema.cue:17:6
+ cue export 200.data.json 200.schema.cue -d '#V3'
hobby: field not allowed:
    ./200.data.json:3:5
    ./200.schema.cue:17:6
+ cue vet 200.data.json 200.schema.cue -d '#V3'
hobby: field not allowed:
    ./200.data.json:3:5
    ./200.schema.cue:17:6
```

