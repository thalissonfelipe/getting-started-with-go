package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return errors.New("error")
	}

	w.balance -= value
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
