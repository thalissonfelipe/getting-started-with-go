package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	}

	verifyError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("waited for an error, but none occurred")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		verifyBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		inititalBalance := Bitcoin(20)
		wallet := Wallet{balance: inititalBalance}
		err := wallet.Withdraw(Bitcoin(100))

		verifyBalance(t, wallet, inititalBalance)
		verifyError(t, err)
	})
}
