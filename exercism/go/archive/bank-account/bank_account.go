// Package account is the package to handle all of our bank accounting needs
package account

import "sync"

// Account is the main struct for our accounting setup
type Account struct {
	sync.RWMutex
	IsOpen bool
	Bal    int64
}

// Open is the primary function to instantiate a new Account struct
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{IsOpen: true, Bal: initialDeposit}
}

// Balance is a method to return the balance if an account is open, otherwise it's !ok
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()

	if a.IsOpen {
		return a.Bal, true
	}

	return 0, false
}

// Close is a method to close off a given account, returning the balance if exists
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.IsOpen {
		a.IsOpen = false
		return a.Bal, true
	}

	return 0, false
}

// Deposit is a method to handle a new amount, adding to an existing Bal if the account is open
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.IsOpen || ((a.Bal + amount) < 0) {
		return 0, false
	}

	a.Bal += amount

	return a.Bal, true
}
