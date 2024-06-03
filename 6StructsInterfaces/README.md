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

Let us see a small set of differences between `Methods and Functions`
| Feature| Method| Function|
|----|----|----|
| Definition| Associated with a struct or interface| Independent entity|
| Access| Called on an object instance| Called directly by its name|
| Implicit Parameter| Object instance (usually named `receiver`) | No implicit parameter|
| Access Control| Public (uppercase) or private (lowercase)| Public (uppercase) or private (lowercase)
| Purpose| Operate on data of the object| Perform general-purpose tasks|