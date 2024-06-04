# Structs 

I already have some experience with structs not only in GO but also in C so we are not going in depth with this.
But for the the sake of it i am going to add the syntax to structs.

```go
type Rectangle struct{
    Length float64
    Breadth float64
}
```

# Methods

For so long we have been using functions in GO, but without knowing we also have been using methods like `t.Errorf()`.

Let us see a small set of differences between **Methods and Functions**
| Feature| Method| Function|
|----|----|----|
| Definition| Associated with a struct or interface| Independent entity|
| Access| Called on an object instance| Called directly by its name|
| Implicit Parameter| Object instance (usually named `receiver`) | No implicit parameter|
| Access Control| Public (uppercase) or private (lowercase)| Public (uppercase) or private (lowercase)
| Purpose| Operate on data of the object| Perform general-purpose tasks|

Syntax for Methods: 
```go
func (receiverName ReceiverType) MethodName(args) returnType {
}
```

# Interfaces

Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety

### Concept of coupling in programming 

- **Highly coupled code**: high coupling in Go occurs when different parts of your code are tightly linked and rely heavily on each other's internal workings. This results in: 
    - Less maintainable code
    - Reusing code becomes a problem 
    - Testing becomes complex
- **Highly decoupled code**: basically when different part of your code are not heavily linked.

# Table driven tests

Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests. 
They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.

To see how easily that is actually try addding a triangle.