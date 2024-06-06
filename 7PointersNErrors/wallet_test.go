package main

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(20))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		// Now here we use format specifiers as %s bcoz we implemented
		// STRINGER Interface with String() method
		t.Errorf("got %s want %s", got, want)
	}
}