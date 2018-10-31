Just working through The Go Programming Language book by Donovan and Kernighan.

:= is short variable delcaration

++ and -- are statements, not expressions, so i++ is valid, but j = i++ is not

_ is the blank identifier

#### Preferred init of strings
##### When the initial value is important
// can't be used for package level variables
s := ""

##### When the default value for the type, "" for strings, is okay
var s string

### Loops

#### for loop
for initialization; condition; post {
    ...
}

#### traditional while loop
for condition {
    ...
}

#### infinite loop
for {

}

#### loop over range
for _, arg := range slice {
    ...
}

### Structs
type person struct {
    name string
    age int
}

joe := person{
    name: "John",
    age: 32,
}

joe.name

