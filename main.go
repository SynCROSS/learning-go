package main

import (
	"fmt"
	"main/accounts"
)

func main() {
	account1 := accounts.CreateAccount("SynCROSS")
	account1.Deposit(1000000000)
	e := account1.Withdraw(1000000000 + 1)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(account1.GetOwner()+"'s Balance:", account1.GetBalance()) // * It Maybe Same Result because of an error.
}
