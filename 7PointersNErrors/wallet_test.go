package main

import "testing"

func TestWallet(t *testing.T) {

	check := func (t testing.TB, wallet Wallet, want Bitcoin)  {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	checkError := func (t testing.TB, got error, want string)  {
		t.Helper()

		if got == nil {
			// This .Fatal will stop the test if it's called
			// This is because we don't want to make any more assertions on the error returned if there isn't one around. 
			// Without this the test would carry on to the next step and panic because of a nil pointer.
			t.Fatal("Didn't get an error but wanted one")
		}

		// This .Error() helps us to convert errors into a string
		if got.Error() != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		check(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		// Previously Withdraw method was just modifying the balance of wallet
		// And was not returning anything so we were trying to assign err to no value
		// This was solved by adding a return value to it
		err := wallet.Withdraw(Bitcoin(100))

		// What values are we passing 
		// we are passing err so that we can send the return msg that we got
		// we are passing the string so check whether this is the string we want or not
		// This .Error() basically helps us pass this error msg as a string
		checkError(t, err, ErrInsufficientFunds.Error())
		check(t, wallet, startingBalance)
	})
}