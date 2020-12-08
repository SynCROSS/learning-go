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