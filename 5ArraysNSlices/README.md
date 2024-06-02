# Arrays and Slices 

As I already have knowledge about arrays and slices so i am not going to discuss about it here.
But there are a few things that should be kept in mind:

- Slices cant be compared directly, so if you wish to compare slices few workarounds are:
    - Write a function that iterates over each and every element
    - Another way is through `reflect.DeepEqual()` which basically compares if two variables are equal or not.
    - With go 1.21 onwards `slices.Equal()` method comes with standard slices package, but for this to work the elements should be comparable.

- Up till now we only know one way to create slices i.e. `var sliceExample []type`, but there are other ways to declare slices.
    - `sum := make([]int, lengthOfSlice)` this type is used when we want to create a slice of given size so that other elements are initialized to zero, and indexed are available.

# Getting to know about varargs 

varargs (short for variable arguments) refer to functions that can accept a variable number of arguments of the same type.

```go
func sum(nums ...int) int {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}

sum(1, 2, 3)  // Returns 6
sum(10)        // Returns 10 (can handle zero arguments)
```
