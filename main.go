package main

import (
	"fmt"
	"main/accounts"
)

func main() {
	account1 := accounts.CreateAccount("SynCROSS")
	fmt.Println(account1) // * ${SynCROSS 0} -> SynCROSS's Current Balance:0
}
