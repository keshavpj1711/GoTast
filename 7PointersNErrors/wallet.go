package main

import "fmt"

type Bitcoin float64

type Wallet struct{
	balance Bitcoin
}

// Implementing a Stringer which is an interface which stores info on  
// how to display the datatype when it's printed
type Stringer interface{
	String() string
}
// This function takes care of what to return when we ask for printing Bitcoin datatype
func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

// This time we are passing a pointer rather than passing a copy of var
func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) {
	w.balance -= amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}