package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// * Go can let developer make the Error
// * Because Go doesn't have try~catch, exception, etc.
// * So Go Developer can make the error variable like below.
var errNoMoney = errors.New("You did NOT have enough money to withdraw")

// CreateAccount creates new Account
func CreateAccount(owner string) *Account {
	newAccount := Account{owner: owner, balance: 0}
	return &newAccount
}

// Deposit for depositing x amount on account
func (a *Account) Deposit(amount int) {
	// * asterisk can edit data directly not using copy.
	// * it is pointer receiver because it's pointer.
	a.balance += amount
}

// GetBalance returns account's balance
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw for withdrawing x amount on account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// * Making custom error can stop behavior in Go.
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
