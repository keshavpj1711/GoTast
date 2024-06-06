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
			t.Fatal("Didn't get an error but wanted one")
		}

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

		checkError(t, err, "can't withdraw insufficient funds")
		check(t, wallet, startingBalance)
	})
}