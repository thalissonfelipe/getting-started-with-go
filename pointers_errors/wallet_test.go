package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		verifyBalance(t, wallet, Bitcoin(10))
		verifyNonexistentError(t, err)
	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		inititalBalance := Bitcoin(20)
		wallet := Wallet{balance: inititalBalance}
		err := wallet.Withdraw(Bitcoin(100))

		verifyBalance(t, wallet, inititalBalance)
		verifyError(t, err, ErrInsufficientBalance)
	})
}

func verifyBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	}
}

func verifyNonexistentError(t *testing.T, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("unexpected error ocurred")
	}
}

func verifyError(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("waited for an error, but none occurred")
	}

	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	}
}
