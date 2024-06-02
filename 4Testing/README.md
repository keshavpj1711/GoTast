# Benchmarking and Examples 

Some other features of GO as a language are Benchmarks and adding Examples.

## Benchmarking

First of all let us look at an example of what is benchmark and how it's written: 

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
```

This was a simple benchmark written in go for our repeat function, the output that we got was: 

```go
> go test -bench=.
goos: linux
goarch: amd64
pkg: 4Testing/iteration
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkRepeat-16       6147052               192.2 ns/op
PASS
ok      4Testing/iteration      1.382s
```

As you can see how much time on an avg does a single operation take i.e. 192.2 ns/op

## Examples 

Examples are a simple way of documenting your code of how does it work.

```go
func ExampleRepeat()  {
	repeat := Repeat("k", 4)
	fmt.Println(repeat)
	// Output: kkkk
}
```

So when you run `go test -v` it gives us whether the example ran fine or failed

