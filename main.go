package main

import (
	"fmt"
	"log"
	"main/accounts"
)

func main() {
	account1 := accounts.CreateAccount("SynCROSS")
	account1.Deposit(1000000000)
	fmt.Println(account1.GetBalance())
	e := account1.Withdraw(1000000000 + 1)
	if e != nil {
		// * Exit program after println
		log.Fatalln(e) // * It Maybe print Error.
	}
	fmt.Println(account1.GetBalance()) // * It Maybe Same Result because of an error.
}
