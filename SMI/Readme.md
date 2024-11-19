## Struct

A `struct` is a sequence of named elements, called fields, each of which has a name and a type.
Fields names may be specified explicitly or implicitly.
Within a struct, non-blank field names must be unique.

```go
// An empty struct
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ floate32 // padding
	A *[]int
	F func()
}
```

## methods

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

```go
func (receiverName ReceiverType) MethodName(args)
```
When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable.
In many other programming languages this is done implicitly, and you access the receiver via `this`

## Interface

```go
type Shape interface {
	Area() float64
}
```