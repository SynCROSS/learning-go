package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// CreateAccount creates new Account
func CreateAccount(owner string) *Account{
  newAccount := Account{owner: owner, balance: 0}
  return &newAccount
}
// Deposit for depositing x amount on account
func (a Account) Deposit(amount int)  {
// * This is how difference between normal function and method.
// * '(receiver_name type_name)' is 'receiver' called in Go.

// * However, there is a rule that 'receiver' 's name
// * should be a lowercase of the first letter of the struct.

// * this sentence can not make balance amount.
// * because Go makes copy of object or struct
// * if Object or Struct is assigned to variable.
// * https://flaviocopes.com/go-copying-structs/
	a.balance += amount
}

// GetBalance returns account's balance
func (a Account) GetBalance() int {
	return a.balance
}