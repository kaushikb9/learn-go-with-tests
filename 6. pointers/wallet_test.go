package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	// t.Run("testing Deposit with pointer notation", func(t *testing.T) {
	// 	wallet := Wallet{}
	// 	walletAddr := &wallet
	// 	walletAddr.Deposit(Bitcoin(10))
	// 	got := walletAddr.Balance()
	// 	fmt.Printf("in Test: wallet = %v\n walletAddr-as-v = %v\n walletAddr-as-p = %p\n", wallet, walletAddr, walletAddr)
	// 	fmt.Printf("in Test: wallet.balance = %v\n &wallet.balance = %v\n (&wallet).balance = %v\n walletAddr.balance = %v\n", wallet.balance, &wallet.balance, (&wallet).balance, walletAddr.balance)
	// 	want := Bitcoin(10)
	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}

	// })

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("testing Deposit directly", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("testing Withdrawal directly", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("testing negative Withdrawals", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
