package main

import (
	"fmt"
	"main/accounts"
)

func main() {
  account1 := accounts.CreateAccount("SynCROSS")
  fmt.Println(account1) // *  &{SynCROSS 0} -- returned Object's address
  // * https://stackoverflow.com/questions/56112289/how-to-print-the-address-of-struct-variable-in-go
}
