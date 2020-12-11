package main

import (
	"fmt"
	"main/accounts"
)

func main() {
	account1 := accounts.CreateAccount("SynCROSS")
	account1.Deposit(1000000000)
  fmt.Println(account1.GetBalance())
  account1.Withdraw(1000000000 + 1)
  fmt.Println(account1.GetBalance()) // * It Maybe Same Result because of an error.
}
