# Pointers 

We are not going to discuss everything in GO about pointers from the basics. 
So here are the things which we should actually be worried about or should i say we should be glad to know.

## No need to deref pointer??

Now you might wonder, why did they pass? We didn't dereference the pointer in the function, like so:

```go
func (w *Wallet) Balance() int {
	return (*w).balance
}
```

and seemingly addressed the object directly. 
In fact, the code above using `(*w)` is absolutely valid. 
However, the makers of Go deemed this notation cumbersome, so the language permits us to write `w.balance`, **without an explicit dereference.** 

>These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
