package hello

hello: "world"

#A: {
    foo: string
}

// to cause a load error, remove the '&'
// to cause a build error, change '#A' to '#B'
// to cause a validation error, change foo to '1'
// note that all values need not be concrete to pass the validation
a: #A & {
    foo: "bar"
}
