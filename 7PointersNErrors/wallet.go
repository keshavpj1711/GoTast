package main

type Wallet struct{
	balance float64
}

func (w Wallet) Deposit(a float64) {
	w.balance += a	
}

func (w Wallet) Balance() float64 {
	return w.balance
}