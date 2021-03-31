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

var ErrInsufficientBalance = errors.New("unable to withdraw: insufficient balance")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= value
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
