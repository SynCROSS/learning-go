package main

import (
	"fmt"
)

func fibonacci(number int) {
	var total, tmp, preTmp int=0,1,1

  for {
    total=preTmp+tmp
    fmt.Println(total)
    preTmp=tmp
    tmp=total

    if total>number{
      break
    }
  }
}

func printElement() {
	fibonacci(18)
}