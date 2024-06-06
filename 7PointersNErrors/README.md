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

### Struct pointers are a thing here...

>These pointers to structs even have their own name: **struct pointers** and they are automatically dereferenced.

# Creating new datatypes from existing ones

Go lets you create new types from existing ones.

The syntax is `type MyName OriginalType`, For example:

```go
type Bitcoin int

type Wallet struct{
    balance Bitcoin
}
```

This will be used everywhere else in the code not just here, so if you are **talking about balance it has to be of datatype Bitcoin** in your code.

So if you are passing balance somewhere else in the code you have to pass it by this, 

```go
wallet.Deposit(Bitcoin(10))
```

> To understand more about it it's better to look at how it's used in the code

# Stringer Interface

As we know we can create new datatypes from existing ones and we can also create different datatypes using structs. 
But one problem stands is how we print these datatypes, bcoz on printing created structs it prints things that we don't want.
This is solved by using **Stringer interface** for that datatype. 

## What is Stringer Interface and how is it used??

This interface defines a single method named `String() string`. 
The purpose of a Stringer is to provide a custom string representation for a type.

```go
type Stringer interface {
	String() string
}

// A Method to print a custom string when we print any Bitcoin datatype
func (b Bitcoin) String() string {
    fmt.Sprintf("%f BTC", b)
}
```

## An Example 

- **Data Type (User):** Let's say you have a **User** struct that stores a username and email address.
- **Stringer (Name Tag Maker):** You can implement the Stringer interface for User.
- **Name Tag (String method):** In the `String()` method, you can define how you want the user information to be displayed. For instance, it could show "`Username: [username] - Email: [email]`".

This way, when you print a User object, you'll see the user's name and email instead of just technical details.
