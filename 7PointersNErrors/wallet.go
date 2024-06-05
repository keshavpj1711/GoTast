package main

type Wallet struct{
	balance float64
}

// This time we are passing a pointer rather than passing a copy of var
func (w *Wallet) Deposit(a float64) {
	w.balance += a	
}

func (w *Wallet) Balance() float64 {
	return w.balance
}