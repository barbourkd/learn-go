package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is a stupid int type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a Wallet you can put money in
type Wallet struct {
	balance Bitcoin
}

// Deposit puts money into the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw pulls money out of the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("attempted to withdraw amount larger than balance")
	}

	w.balance -= amount
	return nil
}

// Balance returns the current wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
