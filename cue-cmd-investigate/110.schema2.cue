package MySchema

#guid: =~"^[A-Za-z0-9]{27}$"

#name: =~ "^Uzuverse ·\u00a0[A-Za-z0-9-]{1,40}$"
// verify that the source file contains all required fields,
// vet the values of the fields,
// and allow for extra fields

#Project: {
    project: {
        guid: #guid
        name: #name
    }
    ... // other fields
}

#Cli: {
    cli: config: string
    ... // other fields
}
