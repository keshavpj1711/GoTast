# Understanding Dependency Injection

It's basically when one of the things in your code depends upon another like in this case, 
our `Greet()` func was initially hard coded with a dependecy that the input will be only of type `bytes.Buffer`.\
But after that we removed that hard coded dependency with `io.Writer` interface as every other thing was implementing this be it:
- `bytes.Buffer` for testing purposes or,
- `os.Stdout` for real world use

> The above is what i understood of DI and not totally correct.

## What we are going to do?

Testing out `fmt.Printf()` is actually a pain reason being it prints to stdout on calling so we need to figure out a way.

## Looking at source code of `fmt.Printf`

```go
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

So below the hood `Printf()` is just calling `Fprintf()`

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

### What's an `io.Writer`?

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

> So basically we are passing `os.Stdout` as `io.Writer`

> So under the covers we're ultimately using Writer to send our greeting somewhere.

